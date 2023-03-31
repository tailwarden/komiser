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
import SelectCheckbox from '../../../select-checkbox/SelectCheckbox';
import Select from '../../../select/Select';
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
    <div className="w-full rounded-lg bg-white py-4 px-6 pb-6">
      <div className=" -mx-6 flex flex-wrap items-center justify-between gap-4 border-b border-black-200/40 px-6 pb-4">
        <div>
          <p className="text-sm font-semibold text-black-900">Cost explorer</p>
          <div className="mt-1"></div>
          <p className="text-xs text-black-300">
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
        {!chartData && (
          <div className="relative flex flex-col items-center">
            <p className="mt-10 text-lg text-black-900">
              No data for this time period
            </p>
            <Image
              src="/assets/img/others/empty-state-cost-explorer.png"
              width={940}
              height={330}
              alt="No data to display image"
              className="absolute top-0"
            />
          </div>
        )}
      </div>
    </div>
  );
}

export default DashboardCostExplorerCard;
