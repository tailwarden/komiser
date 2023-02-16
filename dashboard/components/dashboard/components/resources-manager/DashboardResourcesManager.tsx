import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Doughnut } from 'react-chartjs-2';
import Grid from '../../../grid/Grid';
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
    <div className="w-full rounded-lg bg-white py-4 px-6 pb-6">
      <div className="-mx-6 flex items-center justify-between border-b border-black-200/40 px-6 pb-4">
        <div>
          <p className="text-sm font-semibold text-black-900">Cloud map</p>
          <div className="mt-1"></div>
          <p className="text-xs text-black-300">
            Analyze which regions have active resources
          </p>
        </div>
        <div className="h-[60px]"></div>
      </div>
      <div className="mt-4"></div>
      <Grid gap="sm">
        <Select
          label="Group by"
          value={query}
          options={select.values}
          displayValues={select.displayValues}
          onChange={handleChange}
        />
        <Select
          label="Group by"
          value={query}
          options={select.values}
          displayValues={select.displayValues}
          onChange={handleChange}
        />
      </Grid>
      <div className="mt-4"></div>
      <Doughnut data={chartData} options={options} />
    </div>
  );
}

export default DashboardResourcesManager;
