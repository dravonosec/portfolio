package Java.Logic;

import java.util.List;

class Interpolation {
    private List<Double> getInterpolatedYAxisValues(
        List<Double> uniqueXAxisValues,
        List<Double> xAxisValues,
        List<Double> yAxisValues
    ) {
        double[] origXValues = fixRepeatValues(xAxisValues.stream().mapToDouble(Double::doubleValue).toArray());
        double[] origYValues = yAxisValues.stream().mapToDouble(Double::doubleValue).toArray();
        double[] extended = uniqueXAxisValues.stream().mapToDouble(Double::doubleValue).toArray();

        LinearInterpolator interpolator = new LinearInterpolator();
        PolynomialSplineFunction function = interpolator.interpolate(origXValues, origYValues);

        List<Double> extendedYAxisValues = new ArrayList<>();
        for (double x : extended) {
            if (x < origXValues[0]) {
                // Экстраполяция для значений меньше min(origXValues)
                extendedYAxisValues.add(extrapolateLeft(x, origXValues, origYValues));
            } else if (x > origXValues[origXValues.length - 1]) {
                // Экстраполяция для значений больше max(origXValues)
                extendedYAxisValues.add(extrapolateRight(x, origXValues, origYValues));
            } else {
                // Интерполяция для значений внутри диапазона
                extendedYAxisValues.add(function.value(x));
            }
        }

        return extendedYAxisValues;
    }

    private double extrapolateLeft(double x, double[] origXValues, double[] origYValues) {
        double x0 = origXValues[0];
        double y0 = origYValues[0];
        double x1 = origXValues[1];
        double y1 = origYValues[1];

        // Линейная экстраполяция на основе первых двух точек
        return Math.max(y0 + (y1 - y0) / (x1 - x0) * (x - x0), 0.0);
    }

    private double extrapolateRight(double x, double[] origXValues, double[] origYValues) {
        int lastIndex = origXValues.length - 1;
        double xLast = origXValues[lastIndex];
        double yLast = origYValues[lastIndex];
        double xPrev = origXValues[lastIndex - 1];
        double yPrev = origYValues[lastIndex - 1];

        // Линейная экстраполяция на основе последних двух точек
        return Math.max(yLast + (yLast - yPrev) / (xLast - xPrev) * (x - xLast), 0.0);
    }
}