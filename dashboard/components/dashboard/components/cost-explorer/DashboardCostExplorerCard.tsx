import {
  BarElement,
  CategoryScale,
  Chart as ChartJS,
  Legend,
  LinearScale,
  Title,
  Tooltip
} from 'chart.js';
import Image from 'next/image';
import { Dispatch, SetStateAction } from 'react';
import { Bar } from 'react-chartjs-2';

import Button from '@components/button/Button';
import SelectCheckbox from '@components/select-checkbox/SelectCheckbox';
import Select from '@components/select/Select';
import { CloudIcon } from '@components/icons';
import {
  CostExplorerQueryDateProps,
  CostExplorerQueryGranularityProps,
  CostExplorerQueryGroupProps,
  DashboardCostExplorerData
} from './hooks/useCostExplorer';
import useCostExplorerChart from './hooks/useCostExplorerChart';

ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend
);

type DashboardCostExplorerCardProps = {
  data: DashboardCostExplorerData[] | undefined;
  queryGroup: CostExplorerQueryGroupProps;
  setQueryGroup: Dispatch<SetStateAction<CostExplorerQueryGroupProps>>;
  queryGranularity: CostExplorerQueryGranularityProps;
  setQueryGranularity: Dispatch<
    SetStateAction<CostExplorerQueryGranularityProps>
  >;
  queryDate: CostExplorerQueryDateProps;
  setQueryDate: Dispatch<SetStateAction<CostExplorerQueryDateProps>>;
  exclude: string[];
  setExclude: Dispatch<SetStateAction<string[]>>;
};

function DashboardCostExplorerCard({
  data,
  queryGroup,
  setQueryGroup,
  queryGranularity,
  setQueryGranularity,
  queryDate,
  setQueryDate,
  exclude,
  setExclude
}: DashboardCostExplorerCardProps) {
  const {
    chartData,
    options,
    groupBySelect,
    granularitySelect,
    dateSelect,
    handleFilterChange
  } = useCostExplorerChart({
    data,
    setQueryGroup,
    queryGranularity,
    setQueryGranularity,
    setQueryDate
  });

  return (
    <div className="w-full rounded-lg bg-white px-6 py-4 pb-6">
      <div className="-mx-6 flex flex-wrap items-center justify-between gap-4 border-b border-gray-300 px-6 pb-4">
        <div>
          <p className="text-sm font-semibold text-gray-950">Cost explorer</p>
          <div className="mt-1"></div>
          <p className="text-xs text-gray-500">
            Visualise, understand, and manage your infrastructure costs and
            usage
          </p>
        </div>
      </div>
      <div className="mt-4"></div>
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        <Select
          label="Group by"
          value={queryGroup}
          values={groupBySelect.values}
          displayValues={groupBySelect.displayValues}
          handleChange={newValue => handleFilterChange('group', newValue)}
        />
        <SelectCheckbox
          label="Excluded"
          query={queryGroup}
          exclude={exclude}
          setExclude={setExclude}
        />
        <Select
          label="Granularity"
          value={queryGranularity}
          values={granularitySelect.values}
          displayValues={granularitySelect.displayValues}
          handleChange={newValue => handleFilterChange('granularity', newValue)}
        />
        <Select
          label="Period"
          value={queryDate}
          values={dateSelect.values}
          displayValues={dateSelect.displayValues}
          handleChange={newValue => handleFilterChange('date', newValue)}
        />
      </div>
      <div className="mt-8"></div>
      <div className="h-full min-h-[22rem]">
        {chartData && <Bar data={chartData} options={options} />}
        {/* ⤵ will be removed when cost is supported at Resource level */}
        {queryGroup === 'Resource' && (
          <div className="relative flex flex-col items-center bg-empty-cost-explorer h-[330px] w-full">
            <div className="mt-10 text-lg text-black-900 border border-gray-200 px-8 py-6 flex bg-white">
              <div>
                <p className="text-lg">
                  Cost at resource level not yet supported
                </p>
                <p className="text-sm text-gray-400 mb-4">
                  We recommend our cloud version, Tailwarden, <br />
                  as it supports accurate costs at the resource level
                </p>

                <Button
                  size="sm"
                  gap="md"
                  asLink
                  href="https://tailwarden.com/?utm_source=komiser"
                  target="_blank"
                >
                  <CloudIcon width="24" /> Discover Tailwarden
                </Button>
              </div>
              <Image
                src="/assets/img/purplin/rocket.svg"
                alt="Purplin on a Rocket"
                width="115"
                height="124"
              />
            </div>
          </div>
        )}
        {!chartData && queryGroup !== 'Resource' && (
          <div className="relative flex flex-col items-center bg-empty-cost-explorer h-[330px] w-full">
            <div className="mt-10 text-lg text-black-900 border border-gray-200 px-8 py-6 flex bg-white">
              <div>
                <p className="text-lg">No data for this time period</p>
                <p className="text-sm text-gray-400 mb-4">
                  Our cloud version, Tailwarden, supports <br />
                  historical costs from certain cloud providers
                </p>
                <Button
                  size="sm"
                  gap="md"
                  asLink
                  href="https://tailwarden.com/?utm_source=komiser"
                  target="_blank"
                >
                  <CloudIcon width="24" /> Discover Tailwarden
                </Button>
              </div>
              <Image
                src="/assets/img/purplin/rocket.svg"
                alt="Purplin on a Rocket"
                width="115"
                height="124"
              />
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

export default DashboardCostExplorerCard;
