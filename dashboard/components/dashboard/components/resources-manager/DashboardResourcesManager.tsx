import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Doughnut } from 'react-chartjs-2';
import Select from '../select/Select';
import useResourcesManager from './hooks/useResourcesManager';
import useResourcesManagerChart from './hooks/useResourcesManagerChart';

ChartJS.register(ArcElement, Tooltip, Legend);

function DashboardResourcesManager() {
  const { loading, data, error, fetch } = useResourcesManager();
  const { chartData, options, select, query, handleChange } =
    useResourcesManagerChart({ data });

  if (loading) return <>Loading</>;

  if (error) return <>Error</>;

  return (
    <div className="flex w-full flex-col gap-4 rounded-lg bg-white pl-2 pr-4 pb-6">
      <div className="border-white-150 -ml-2 -mr-4 flex flex-wrap items-center justify-between border-b">
        <div className="px-6 py-4">
          <p className="text-sm font-semibold text-black-900">
            Monthly cost break down
          </p>
          <div className="mt-1"></div>
          <p className="text-xs text-black-300">Top 5 biggest costs</p>
        </div>
        <div className="w-full px-6 py-4 md:w-auto md:flex-row">
          <Select
            label="Group by"
            value={query}
            options={select.values}
            displayValues={select.displayValues}
            onChange={handleChange}
          />
        </div>
      </div>
      <Doughnut data={chartData} options={options} />
    </div>
  );
}

export default DashboardResourcesManager;
