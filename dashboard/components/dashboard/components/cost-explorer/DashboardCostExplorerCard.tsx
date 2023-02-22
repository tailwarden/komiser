import {
  BarElement,
  CategoryScale,
  Chart as ChartJS,
  Legend,
  LinearScale,
  Title,
  Tooltip
} from 'chart.js';
import { Dispatch, SetStateAction } from 'react';
import { Bar } from 'react-chartjs-2';
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
  data: DashboardCostExplorerData | undefined;
  queryGroup: CostExplorerQueryGroupProps;
  setQueryGroup: Dispatch<SetStateAction<CostExplorerQueryGroupProps>>;
  queryGranularity: CostExplorerQueryGranularityProps;
  setQueryGranularity: Dispatch<
    SetStateAction<CostExplorerQueryGranularityProps>
  >;
  queryDate: CostExplorerQueryDateProps;
  setQueryDate: Dispatch<SetStateAction<CostExplorerQueryDateProps>>;
};

function DashboardCostExplorerCard({
  data,
  queryGroup,
  setQueryGroup,
  queryGranularity,
  setQueryGranularity,
  queryDate,
  setQueryDate
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
        <div className="flex w-full flex-col gap-4 md:w-auto md:flex-row">
          <Select
            label="Group by"
            values={groupBySelect.values}
            value={queryGroup}
            displayValues={groupBySelect.displayValues}
            onChange={e => handleFilterChange(e, 'group')}
          />
          <Select
            label="Granularity"
            values={granularitySelect.values}
            value={queryGranularity}
            displayValues={granularitySelect.displayValues}
            onChange={e => handleFilterChange(e, 'granularity')}
          />
          <Select
            label="Period"
            values={dateSelect.values}
            value={queryDate}
            displayValues={dateSelect.displayValues}
            onChange={e => handleFilterChange(e, 'date')}
          />
        </div>
      </div>
      <div className="mt-8"></div>
      <div className="h-full min-h-[22rem]">
        {chartData && <Bar data={chartData} options={options} />}
      </div>
    </div>
  );
}

export default DashboardCostExplorerCard;
