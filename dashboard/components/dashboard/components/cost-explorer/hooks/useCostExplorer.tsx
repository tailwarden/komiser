import { useEffect, useState } from 'react';
import settingsService from '../../../../../services/settingsService';
import dateHelper, {
  lastMonth,
  lastSixMonths,
  lastThreeMonths,
  lastTwelveMonths
} from '../utils/dateHelper';

export type DashboardCostExplorerData = {
  date: string;
  datapoints: {
    name: string;
    amount: number;
  }[];
};

export type CostExplorerQueryGroupProps =
  | 'provider'
  | 'service'
  | 'region'
  | 'account'
  | 'view';
export type CostExplorerQueryGranularityProps = 'monthly' | 'daily';
export type CostExplorerQueryDateProps =
  | 'lastMonth'
  | 'lastThreeMonths'
  | 'lastSixMonths'
  | 'lastTwelveMonths';

function useCostExplorer() {
  const [loading, setLoading] = useState(true);
  const [data, setData] = useState<DashboardCostExplorerData[]>();
  const [error, setError] = useState(false);
  const [queryGroup, setQueryGroup] =
    useState<CostExplorerQueryGroupProps>('provider');
  const [queryGranularity, setQueryGranularity] =
    useState<CostExplorerQueryGranularityProps>('monthly');
  const [queryDate, setQueryDate] =
    useState<CostExplorerQueryDateProps>('lastSixMonths');
  const [exclude, setExclude] = useState<string[]>([]);

  function fetch(
    group: CostExplorerQueryGroupProps = 'provider',
    newGranularity: CostExplorerQueryGranularityProps = 'monthly',
    start: string = dateHelper.getLastSixMonths(),
    end: string = dateHelper.getToday()
  ) {
    if (!loading) {
      setLoading(true);
    }

    if (error) {
      setError(false);
    }

    const granularity = newGranularity.toUpperCase();
    const payload = {
      group,
      granularity,
      start,
      end,
      exclude
    };
    const payloadJson = JSON.stringify(payload);

    settingsService.getCostExplorer(payloadJson).then(res => {
      if (res === Error) {
        setLoading(false);
        setError(true);
      } else {
        setLoading(false);
        setData(res);
      }
    });
  }

  useEffect(() => {
    let startDate = '';
    let endDate = '';

    if (queryDate === 'lastMonth') {
      [startDate, endDate] = lastMonth;
    }
    if (queryDate === 'lastThreeMonths') {
      [startDate, endDate] = lastThreeMonths;
    }
    if (queryDate === 'lastSixMonths') {
      [startDate, endDate] = lastSixMonths;
    }
    if (queryDate === 'lastTwelveMonths') {
      [startDate, endDate] = lastTwelveMonths;
    }

    fetch(queryGroup, queryGranularity, startDate, endDate);
  }, [queryGroup, queryGranularity, queryDate, exclude]);

  return {
    loading,
    data,
    error,
    fetch,
    queryGroup,
    setQueryGroup,
    queryGranularity,
    setQueryGranularity,
    queryDate,
    setQueryDate
  };
}

export default useCostExplorer;
