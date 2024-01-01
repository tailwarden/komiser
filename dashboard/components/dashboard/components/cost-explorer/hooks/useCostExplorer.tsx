import { useEffect, useRef, useState } from 'react';
import settingsService from '../../../../../services/settingsService';
import dateHelper, {
  thisMonth,
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
  | 'view'
  | 'Resource';
export type CostExplorerQueryGranularityProps = 'monthly' | 'daily';
export type CostExplorerQueryDateProps =
  | 'thisMonth'
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
  const previousQueryGroup = useRef(queryGroup);

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

    let startDate = '';
    let endDate = '';

    if (queryDate === 'thisMonth') {
      [startDate, endDate] = thisMonth;
    }
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

    const granularity = queryGranularity.toUpperCase();
    const payload = {
      group: queryGroup,
      granularity,
      start: startDate,
      end: endDate,
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
    if (queryGroup !== previousQueryGroup.current) {
      setExclude([]);
    }
    previousQueryGroup.current = queryGroup;
    fetch();
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
    setQueryDate,
    exclude,
    setExclude
  };
}

export default useCostExplorer;
