import { ChartData, ChartOptions } from 'chart.js';
import { ChangeEvent, Dispatch, SetStateAction } from 'react';
import formatNumber from '../../../../../utils/formatNumber';
import {
  ResourcesManagerData,
  ResourcesManagerGroupBySelectProps,
  ResourcesManagerQuery
} from './useResourcesManager';

export type ResourcesManagerChartProps = {
  name: string;
  amount: number;
};

type useResourcesManagerChartProps = {
  data: ResourcesManagerData | undefined;
  setQuery: Dispatch<SetStateAction<ResourcesManagerQuery>>;
};

function useResourcesManagerChart({
  data,
  setQuery
}: useResourcesManagerChartProps) {
  function handleChange(e: ChangeEvent<HTMLSelectElement>) {
    setQuery(e.currentTarget.value as ResourcesManagerQuery);
  }

  const select: ResourcesManagerGroupBySelectProps = {
    values: ['provider', 'service', 'region', 'account', 'view'],
    displayValues: [
      'Cloud provider',
      'Cloud service',
      'Cloud region',
      'Cloud account',
      'Custom views'
    ]
  };

  const colors = ['#80AAF2', '#F19B6E', '#FBC864', '#9BD6CC', '#B8D987'];

  const sortByDescendingCosts = data?.sort(
    (a: ResourcesManagerChartProps, b: ResourcesManagerChartProps) =>
      b.amount - a.amount
  );

  const chartData: ChartData<'doughnut'> = {
    labels: sortByDescendingCosts?.map(item => item.name),
    datasets: [
      {
        data: sortByDescendingCosts?.map(item =>
          Number(formatNumber(item.amount))
        ) as number[],
        backgroundColor: colors,
        borderColor: '#FFFFFF',
        borderWidth: 3,
        hoverOffset: 15
      }
    ]
  };

  const options: ChartOptions<'doughnut'> = {
    aspectRatio: 2,
    layout: {
      padding: 5
    },
    plugins: {
      legend: {
        position: 'right',
        align: 'center',
        labels: {
          font: {
            family: 'Noto Sans'
          },
          usePointStyle: true,
          padding: 16,
          generateLabels: chart => {
            const dataset = chart.data.datasets;
            const background = dataset[0].backgroundColor as string[];
            return dataset[0].data.map((dataSet, i) => ({
              text: `${' '} ${chart.data.labels![i]} - ${dataSet} ${
                dataSet === 1 ? 'resource' : 'resources'
              }`,
              fontColor: '#091126',
              fillStyle: background[i],
              strokeStyle: '#FFFFFF',
              hidden: !chart.getDataVisibility(i),
              index: i
            }));
          }
        }
      },
      tooltip: {
        backgroundColor: 'rgba(0,0,0,.75)',
        multiKeyBackground: '#282828',
        boxPadding: 8,
        padding: 12,
        usePointStyle: true,
        bodyFont: {
          family: 'Noto Sans'
        },
        callbacks: {
          title: () => '',
          label(label) {
            return `${label.label} - ${label.formattedValue} ${
              Number(label.formattedValue) === 1 ? 'resource' : 'resources'
            }`;
          }
        }
      }
    }
  };

  return { chartData, options, select, handleChange };
}

export default useResourcesManagerChart;
