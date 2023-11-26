import { ArcElement, Chart as ChartJS, Legend, Tooltip } from 'chart.js';
import Image from 'next/image';
import { Dispatch, SetStateAction } from 'react';
import { Doughnut } from 'react-chartjs-2';
import SelectCheckbox from '../../../select-checkbox/SelectCheckbox';
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
  exclude: string[];
  setExclude: Dispatch<SetStateAction<string[]>>;
};

function DashboardResourcesManagerCard({
  data,
  query,
  setQuery,
  exclude,
  setExclude
}: DashboardResourcesManagerCardProps) {
  const { chartData, options, select, handleChange } = useResourcesManagerChart(
    { data, setQuery, initialQuery: query }
  );

  const displayChart =
    chartData && chartData.labels && chartData.labels.length > 0;

  return (
    <div className="w-full rounded-lg bg-white px-6 py-4 pb-6">
      <div className="-mx-6 flex items-center justify-between border-b border-gray-300 px-6 pb-4">
        <div>
          <p className="text-sm font-semibold text-gray-950">
            Resources manager
          </p>
          <div className="mt-1"></div>
          <p className="text-xs text-gray-500">
            Uncover how your resources are distributed
          </p>
        </div>
        <div className="h-[60px]"></div>
      </div>
      <div className="mt-4"></div>
      <div className="grid gap-4 md:grid-cols-2">
        <Select
          label="Group by"
          value={query}
          values={select.values}
          displayValues={select.displayValues}
          handleChange={handleChange}
        />
        <SelectCheckbox
          label="Exclude"
          setExclude={setExclude}
          exclude={exclude}
          query={query}
        />
      </div>
      <div className="mt-4"></div>
      <div>
        {displayChart && <Doughnut data={chartData} options={options} />}
        {!displayChart && (
          <div className="relative flex items-center gap-16 px-4 pt-10">
            <Image
              src="/assets/img/others/empty-state-resources-manager.png"
              width={200}
              height={200}
              alt="No data to display image"
            />
            <p className="text-center text-lg text-gray-950">
              No resource data
              <br />
              to display
            </p>
          </div>
        )}
      </div>
    </div>
  );
}

export default DashboardResourcesManagerCard;
