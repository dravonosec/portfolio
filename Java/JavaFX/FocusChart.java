package Java.JavaFX;

public class FocusChart extends AnchorPane {

    private static final int HEIGHT = 200;
    private static final int COUNT = 30;

    private final AxCvCameraFocusController focusController;

    private final DoubleDataSet focusData;
    private final DoubleDataSet labelData;

    private int maxIndex;


    public FocusChart(AxCvCameraFocusController focusController) {
        this.focusController = focusController;
        maxIndex = 0;

        DefaultNumericAxis xAxis1 = new DefaultNumericAxis(null, null);
        xAxis1.set(0, COUNT);
        xAxis1.setTickUnit(0);
        xAxis1.setTickLabelFormatter(null);
        xAxis1.setMinorTickCount(0);
        xAxis1.setMaxMajorTickLabelCount(0);
        xAxis1.setMinorTickLength(0);
        xAxis1.setAutoRangeRounding(false);
        xAxis1.setAutoUnitScaling(false);
        xAxis1.setAutoGrowRanging(false);
        xAxis1.setAutoRanging(false);

        DefaultNumericAxis yAxis1 = new DefaultNumericAxis(null, null);
        yAxis1.setAutoRangeRounding(false);
        yAxis1.setAutoUnitScaling(true);
        yAxis1.setTickUnit(1.0);
        yAxis1.setAutoRangePadding(0.1);
        yAxis1.setForceZeroInRange(false);
        yAxis1.setAutoGrowRanging(true);
        yAxis1.setAutoRanging(true);

        focusData = new DoubleDataSet("Значение фокуса");

        ErrorDataSetRenderer errorDataSetRenderer = new ErrorDataSetRenderer();
        errorDataSetRenderer.getDatasets().add(focusData);

        labelData = new DoubleDataSet("");
        labelData.add(0.0, 0.0, " ");
        labelData.setStyle(
                DataSetStyleBuilder.newInstance().reset()
                        .setLineColor("blue")
                        .setMarkerColor("blue")
                        .setLineDashes(3, 5, 8, 5)
                        .build()
        );

        LabelledMarkerRenderer labeledMarkerRender = new LabelledMarkerRenderer();
        labeledMarkerRender.enableHorizontalMarker(true);
        labeledMarkerRender.enableVerticalMarker(false);
        AxFxThread.run(() -> labeledMarkerRender.getDatasets().add(labelData));

        CompletableFuture<XYChart> chartFuture = new CompletableFuture<>();
        AxFxThread.run(() -> {
            XYChart chart = new XYChart(xAxis1, yAxis1);
            chart.setAnimated(false);
            chart.getRenderers().clear();
            chart.getRenderers().add(labeledMarkerRender);
            chart.getRenderers().add(errorDataSetRenderer);
            chart.setLegendVisible(false);
            chartFuture.complete(chart);
        });
        XYChart chart = chartFuture.join();

        this.setMaxHeight(HEIGHT);
        this.setMinHeight(HEIGHT);
        this.getChildren().add(AxFxAnchorCons.of(chart).anchor(new AxFxAnchors(0, -15, 0, -20)).set());
    }


    @Override
    public void load() { }

    @Override
    public void reset() { }

    @Override
    public void event(String event) { }

    public void update() {
        showFocus();
    }

    private void showFocus() {
        int focusValue = focusController.getValue();

        AxFxThread.run(() -> {
            if (focusData.getDataCount() > COUNT - 1) {
                focusData.remove(0);
                focusData.add(COUNT - 1, focusValue);
                for (int i = 0; i < focusData.getDataCount(); i++) {
                    focusData.getValues(0)[i]--;
                }
                maxIndex--;
            } else {
                focusData.add(focusData.getDataCount(), focusValue);
            }
            maxIndex = getMaxIndex(focusValue);
            labelData.getYValues()[0] = focusData.getValues(1)[maxIndex];
        });
    }

    private int getMaxIndex(int newFocusValue) {
        double[] values = focusData.getValues(1);

        int targetIndex = 0;
        if (maxIndex == -1) {
            double maxValue = 0.0;

            for (int i = 0; i < focusData.getDataCount(); i++) {
                if (values[i] > maxValue) {
                    maxValue = values[i];
                    targetIndex = i;
                }
            }
            return targetIndex;
        }

        if (values[maxIndex] < newFocusValue) {
            return focusData.getDataCount() - 1;
        } else {
            return maxIndex;
        }
    }

}
