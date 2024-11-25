package Java.Logic;

import org.apache.commons.math3.analysis.ParametricUnivariateFunction;
import org.apache.commons.math3.fitting.AbstractCurveFitter;
import org.apache.commons.math3.fitting.WeightedObservedPoint;
import org.apache.commons.math3.fitting.leastsquares.LeastSquaresBuilder;
import org.apache.commons.math3.fitting.leastsquares.LeastSquaresProblem;
import org.apache.commons.math3.fitting.leastsquares.ParameterValidator;
import org.apache.commons.math3.linear.ArrayRealVector;
import org.apache.commons.math3.linear.DiagonalMatrix;
import org.apache.commons.math3.linear.RealVector;

import java.util.Collection;

@Deprecated
public class Fitter extends AbstractCurveFitter {

    private double[] start;
    private double[] lowerBounds;
    private double[] upperBounds;

    public Fitter() {
        this.start = new double[]{100, 100, 0, 50, 0.5};
        this.lowerBounds = new double[]{1, 20, -2, 0.5, 0};
        this.upperBounds = new double[]{200, 200, 2, 160, 1};
    }

    public Fitter(double[] start, double[] lowerBounds, double[] upperBounds) {
        this.start = start;
        this.lowerBounds = lowerBounds;
        this.upperBounds = upperBounds;
    }

    @Override
    protected LeastSquaresProblem getProblem(Collection<WeightedObservedPoint> observations) {
        final int len = observations.size();
        final double[] target = new double[len];
        final double[] weights = new double[len];

        int count = 0;
        for (WeightedObservedPoint obs : observations) {
            target[count] = obs.getY();
            weights[count] = obs.getWeight();
            count++;
        }

        final AbstractCurveFitter.TheoreticalValuesFunction model = new AbstractCurveFitter.TheoreticalValuesFunction(new MyFunction(), observations);

        return new LeastSquaresBuilder().
                maxEvaluations(Integer.MAX_VALUE).
                maxIterations(Integer.MAX_VALUE).
                start(start).
                target(target).
                weight(new DiagonalMatrix(weights)).
                model(model.getModelFunction(), model.getModelFunctionJacobian()).
                parameterValidator(new MyParameterValidator(lowerBounds, upperBounds)).
                build();
    }

    private static class MyFunction implements ParametricUnivariateFunction {
        @Override
        public double value(double t, double... parameters) {
            double a0 = parameters[0];
            double a1 = parameters[1];
            double a2 = parameters[2];
            double a3 = parameters[3];
            double a4 = parameters[4];
            
            // Функция фиттинга
            return Math.abs(a0 * Math.exp(-t / a1) * Math.sin((t - a2) * 2 * Math.PI * 1 / (4 * a3)) + a4);
        }

        @Override
        public double[] gradient(double t, double... parameters) {
            double a0 = parameters[0];
            double a1 = parameters[1];
            double a2 = parameters[2];
            double a3 = parameters[3];
            double a4 = parameters[4];

            double expTerm = Math.exp(-t / a1);
            double sinTerm = Math.sin((t - a2) * 2 * Math.PI * 1 / (4 * a3));
            double cosTerm = Math.cos((t - a2) * 2 * Math.PI * 1 / (4 * a3));

            double innerTerm = a0 * expTerm * sinTerm + a4;
            double sign = Math.signum(innerTerm);

            // Частные производные по каждому параметру
            double da0 = sign * expTerm * sinTerm;
            double da1 = sign * a0 * t / (a1 * a1) * expTerm * sinTerm;
            double da2 = -sign * a0 * expTerm * cosTerm * 2 * Math.PI * 1 / (4 * a3);
            double da3 = sign * a0 * expTerm * cosTerm * 2 * Math.PI * (t - a2) / (4 * a3 * a3);
            double da4 = sign;

            return new double[]{da0, da1, da2, da3, da4};
        }
    }

    private static class MyParameterValidator implements ParameterValidator {
        private final double[] lowerBounds;
        private final double[] upperBounds;

        public MyParameterValidator(double[] lowerBounds, double[] upperBounds) {
            this.lowerBounds = lowerBounds;
            this.upperBounds = upperBounds;
        }

        @Override
        public RealVector validate(RealVector params) {
            double[] paramArray = params.toArray();
            for (int i = 0; i < paramArray.length; i++) {
                if (paramArray[i] < lowerBounds[i]) {
                    paramArray[i] = lowerBounds[i];
                } else if (paramArray[i] > upperBounds[i]) {
                    paramArray[i] = upperBounds[i];
                }
            }
            return new ArrayRealVector(paramArray);
        }
    }

    public double[] getStart() {
        return start;
    }

    public void setStart(double[] start) {
        this.start = start;
    }

    public double[] getLowerBounds() {
        return lowerBounds;
    }

    public void setLowerBounds(double[] lowerBounds) {
        this.lowerBounds = lowerBounds;
    }

    public double[] getUpperBounds() {
        return upperBounds;
    }

    public void setUpperBounds(double[] upperBounds) {
        this.upperBounds = upperBounds;
    }
}
