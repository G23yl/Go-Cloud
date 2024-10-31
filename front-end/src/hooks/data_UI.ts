// vue-data-ui的配置文件

import { reactive } from "vue"
import type { VueUiDonutConfig, VueUiWheelConfig } from "vue-data-ui"

export const useDataUI = () => {
  // 饼状图的config
  const donutConfig = reactive<VueUiDonutConfig>({
    responsive: false,
    theme: "",
    customPalette: [],
    useCssAnimation: true,
    useBlurOnHover: true,
    userOptions: {
      show: false,
      buttons: {
        tooltip: false,
        pdf: false,
        csv: false,
        img: false,
        table: false,
        labels: false,
        fullscreen: false,
        sort: false,
        stack: false,
        animation: false,
      },
      buttonTitles: {
        open: "Open options",
        close: "Close options",
        tooltip: "Toggle tooltip",
        pdf: "Download PDF",
        csv: "Download CSV",
        img: "Download PNG",
        table: "Toggle table",
        labels: "Toggle labels",
        fullscreen: "Toggle fullscreen",
      },
    },
    translations: { total: "Total", average: "Average" },
    table: {
      show: false,
      responsiveBreakpoint: 400,
      th: { backgroundColor: "#FFFFFF", color: "#1A1A1A", outline: "none" },
      td: {
        backgroundColor: "#FFFFFF",
        color: "#1A1A1A",
        outline: "none",
        roundingValue: 0,
        roundingPercentage: 0,
      },
      columnNames: { series: "Series", value: "Value", percentage: "Percentage" },
    },
    style: {
      fontFamily: "inherit",
      chart: {
        useGradient: true,
        gradientIntensity: 40,
        backgroundColor: "#eeeeee",
        color: "#1a1a1a",
        layout: {
          labels: {
            dataLabels: {
              show: true,
              useLabelSlots: false,
              hideUnderValue: 3,
              prefix: "",
              suffix: "",
            },
            value: { rounding: 0, show: true, formatter: null },
            percentage: {
              color: "#1A1A1A",
              bold: true,
              fontSize: 18,
              rounding: 0,
              formatter: null,
            },
            name: { color: "#1A1A1A", bold: false, fontSize: 14 },
            hollow: {
              show: true,
              total: {
                show: true,
                bold: false,
                fontSize: 18,
                color: "#AAAAAA",
                text: "Total",
                offsetY: 0,
                value: {
                  color: "#1A1A1A",
                  fontSize: 18,
                  bold: true,
                  suffix: "",
                  prefix: "",
                  offsetY: 0,
                  rounding: 0,
                  formatter: null,
                },
              },
              average: {
                show: false,
                bold: false,
                fontSize: 18,
                color: "#AAAAAA",
                text: "",
                offsetY: 0,
                value: {
                  color: "#1A1A1A",
                  fontSize: 18,
                  bold: true,
                  suffix: "",
                  prefix: "",
                  offsetY: 0,
                  rounding: 0,
                  formatter: null,
                },
              },
            },
          },
          donut: { strokeWidth: 64, borderWidth: 1, useShadow: false, shadowColor: "#1A1A1A" },
        },
        comments: { show: true, showInTooltip: true, width: 100, offsetY: 0, offsetX: 0 },
        legend: {
          show: true,
          bold: false,
          backgroundColor: "#eeeeee",
          color: "#1A1A1A",
          fontSize: 13,
          roundingValue: 0,
          roundingPercentage: 0,
        },
        tooltip: {
          show: false,
          color: "#1A1A1A",
          backgroundColor: "#FFFFFF",
          fontSize: 14,
          customFormat: null,
          borderRadius: 4,
          borderColor: "#e1e5e8",
          borderWidth: 1,
          backgroundOpacity: 100,
          position: "center",
          offsetY: 24,
          showValue: true,
          showPercentage: true,
          roundingValue: 0,
          roundingPercentage: 0,
        },
        title: {
          text: "",
          color: "#1A1A1A",
          fontSize: 20,
          bold: true,
          textAlign: "center",
          paddingLeft: 0,
          paddingRight: 0,
          subtitle: { color: "#A1A1A1", text: "", fontSize: 16, bold: false },
        },
      },
    },
  })

  // 轮型图的config
  const wheelConfig = reactive<VueUiWheelConfig>({
    responsive: false,
    theme: "",
    style: {
      fontFamily: "inherit",
      chart: {
        backgroundColor: "#eeeeee",
        color: "#1A1A1A",
        animation: { use: true, speed: 0.5, acceleration: 1 },
        layout: {
          wheel: {
            ticks: {
              rounded: true,
              inactiveColor: "#e1e5e8",
              activeColor: "#6376DD",
              gradient: { show: true, shiftHueIntensity: 10 },
            },
          },
          innerCircle: { show: true, stroke: "#e1e5e8", strokeWidth: 1 },
          percentage: { show: true, fontSize: 48, rounding: 0, bold: true, formatter: null },
        },
        title: {
          text: "",
          color: "#1A1A1A",
          fontSize: 20,
          bold: true,
          textAlign: "center",
          paddingLeft: 0,
          paddingRight: 0,
          subtitle: { color: "#A1A1A1", text: "", fontSize: 16, bold: false },
        },
      },
    },
    userOptions: {
      show: false,
      buttons: {
        tooltip: false,
        pdf: false,
        csv: false,
        img: false,
        table: false,
        labels: false,
        fullscreen: false,
        sort: false,
        stack: false,
        animation: false,
      },
      buttonTitles: {
        open: "Open options",
        close: "Close options",
        pdf: "Download PDF",
        img: "Download PNG",
        fullscreen: "Toggle fullscreen",
      },
    },
  })

  return {
    donutConfig,
    wheelConfig,
  }
}
