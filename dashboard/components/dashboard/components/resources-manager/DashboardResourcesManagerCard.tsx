import React, { Dispatch, SetStateAction } from 'react';
import { Doughnut } from 'react-chartjs-2';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import Select from '../select/Select';
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
};

function DashboardResourcesManagerCard({
  data,
  query,
  setQuery
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
      <Select
        label="Group by"
        value={query}
        options={select.values}
        displayValues={select.displayValues}
        onChange={handleChange}
      />
      <div className="mt-4"></div>
      <Doughnut data={chartData} options={options} />
    </div>
  );
}

export default DashboardResourcesManagerCard;
