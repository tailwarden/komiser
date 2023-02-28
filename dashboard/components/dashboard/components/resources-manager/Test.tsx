import React from 'react';
import { Doughnut } from 'react-chartjs-2';
import ReactDOMServer from 'react-dom/server';

interface TooltipProps {
  label: string;
  value: number;
}

const CustomTooltip: React.FC<TooltipProps> = ({ label, value }) => (
  <div>
    <p>{label}</p>
    <p>{value}</p>
  </div>
);

const data = {
  labels: ['Red', 'Blue', 'Yellow'],
  datasets: [
    {
      label: 'My Dataset',
      data: [10, 20, 30],
      backgroundColor: ['red', 'blue', 'yellow']
    }
  ]
};

const options = {
  plugins: {
    tooltip: {
      enabled: false,
      custom(tooltipModel: any) {
        // `this` is the chart instance
        const tooltipEl = document.getElementById('chartjs-tooltip');
        if (tooltipEl) {
          tooltipEl.style.opacity = '1';
          tooltipEl.style.position = 'absolute';
          tooltipEl.style.top = `${tooltipModel.caretY}px`;
          tooltipEl.style.left = `${tooltipModel.caretX}px`;
          tooltipEl.style.pointerEvents = 'none';
          const tooltipData = {
            label: data.labels[tooltipModel.index],
            value: data.datasets[0].data[tooltipModel.index]
          };
          const tooltipContent = ReactDOMServer.renderToString(
            <CustomTooltip
              label={tooltipData.label}
              value={tooltipData.value}
            />
          );
          tooltipEl.innerHTML = tooltipContent;
        }
      }
    }
  }
};

const MyChart = () => (
  <div style={{ position: 'relative', width: '300px', height: '300px' }}>
    <div id="chartjs-tooltip" className="h-2 border-error-600">
      a
    </div>
    <Doughnut data={data} options={options} />
  </div>
);

export default MyChart;
