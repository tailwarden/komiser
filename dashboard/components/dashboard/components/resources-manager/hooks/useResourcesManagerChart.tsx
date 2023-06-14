import { ChartData, ChartOptions } from 'chart.js';
import { Dispatch, SetStateAction, useState } from 'react';
import { useRouter } from 'next/router';
import formatNumber from '../../../../../utils/formatNumber';
import {
  ResourcesManagerData,
  ResourcesManagerGroupBySelectProps,
  ResourcesManagerQuery
} from './useResourcesManager';

export type ResourcesManagerChartProps = {
  label: string;
  total: number;
};

type useResourcesManagerChartProps = {
  data: ResourcesManagerData | undefined;
  setQuery: Dispatch<SetStateAction<ResourcesManagerQuery>>;
  initialQuery: ResourcesManagerQuery;
};

function useResourcesManagerChart({
  data,
  setQuery,
  initialQuery
}: useResourcesManagerChartProps) {
  const router = useRouter();
  const colors = ['#FF9A87', '#6DB7FF', '#B6D3B4', '#FFB459', '#59ACAC'];

  /* To be un-commented when 'view' is supported 
    const select: ResourcesManagerGroupBySelectProps = {
    values: ['provider', 'service', 'region', 'account', 'view'],
    displayValues: [
      'Cloud provider',
      'Cloud service',
      'Cloud region',
      'Cloud account',
      'Custom views'
    ]
  }; */

  const [currentQuery, setCurrentQuery] = useState(initialQuery);

  const select: ResourcesManagerGroupBySelectProps = {
    values: ['provider', 'service', 'region', 'account'],
    displayValues: [
      'Cloud provider',
      'Cloud service',
      'Cloud region',
      'Cloud account'
    ]
  };

  const sortByDescendingCosts = data?.sort(
    (a: ResourcesManagerChartProps, b: ResourcesManagerChartProps) =>
      b.total - a.total
  );

  const chartData: ChartData<'doughnut'> = {
    labels: sortByDescendingCosts?.map(item => item.label),
    datasets: [
      {
        data: sortByDescendingCosts?.map(item => item.total) as number[],
        backgroundColor: colors,
        borderColor: '#FFFFFF',
        borderWidth: 3,
        hoverOffset: 15
      }
    ]
  };

  const resources = data && data.map(resource => resource.total);
  const sumOfResources =
    resources && resources.reduce((resource, a) => resource + a, 0);

  const options: ChartOptions<'doughnut'> = {
    onClick(_, legendItem, legend) {
      const clickedIndex = legendItem[0].index || 0;
      const labels = legend.config.data.labels ?? [];

      router.push(`/inventory?${currentQuery}:IS:${labels[clickedIndex]}`);
    },
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
            if (sumOfResources) {
              return `${label.label} - ${label.formattedValue} ${
                Number(label.formattedValue) === 1 ? `resource` : `resources`
              } (${((Number(label.raw) / sumOfResources) * 100).toFixed(1)}%)`;
            }
            return undefined;
          }
        }
      }
    }
  };

  function handleChange(newValue: string) {
    setCurrentQuery(newValue as ResourcesManagerQuery);
    setQuery(newValue as ResourcesManagerQuery);
  }

  return { chartData, options, select, handleChange };
}

export default useResourcesManagerChart;
