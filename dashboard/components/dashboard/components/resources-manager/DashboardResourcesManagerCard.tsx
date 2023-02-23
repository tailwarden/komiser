import { ArcElement, Chart as ChartJS, Legend, Tooltip } from 'chart.js';
import { Dispatch, SetStateAction } from 'react';
import { Doughnut } from 'react-chartjs-2';
import CheckboxSelect from '../../../checkbox-select/CheckboxSelect';
import Grid from '../../../grid/Grid';
import Select from '../../../select/Select';
import {
  ResourcesManagerData,
  ResourcesManagerQuery
} from './hooks/useResourcesManager';
import useResourcesManagerChart from './hooks/useResourcesManagerChart';

ChartJS.register(ArcElement, Tooltip, Legend);

type DashboardResourcesManagerCardProps = {
  data: ResourcesManagerData | undefined;
  query: ResourcesManagerQuery;
  setQuery: Dispatch<SetStateAction<ResourcesManagerQuery>>;
  listOfResources: string[];
  exclude: string[];
  setExclude: Dispatch<SetStateAction<string[]>>;
};

function DashboardResourcesManagerCard({
  data,
  query,
  setQuery,
  listOfResources,
  exclude,
  setExclude
}: DashboardResourcesManagerCardProps) {
  const { chartData, options, select, handleChange } = useResourcesManagerChart(
    { data, setQuery }
  );

  return (
    <div className="w-full rounded-lg bg-white py-4 px-6 pb-6">
      <div className="-mx-6 flex items-center justify-between border-b border-black-200/40 px-6 pb-4">
        <div>
          <p className="text-sm font-semibold text-black-900">
            Resources manager
          </p>
          <div className="mt-1"></div>
          <p className="text-xs text-black-300">
            Uncover how your resources are distributed
          </p>
        </div>
        <div className="h-[60px]"></div>
      </div>
      <div className="mt-4"></div>
      <Grid gap="sm">
        <Select
          label="Group by"
          value={query}
          values={select.values}
          displayValues={select.displayValues}
          onChange={handleChange}
        />
        <CheckboxSelect
          label="Exclude"
          listOfResources={listOfResources}
          setExclude={setExclude}
          exclude={exclude}
        />
      </Grid>
      <div className="mt-4"></div>
      <Doughnut data={chartData} options={options} />
    </div>
  );
}

export default DashboardResourcesManagerCard;
