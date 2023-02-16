import { useEffect, useState } from 'react';
import mockDataForDashboard from '../../../utils/mockDataForDashboard';

type Data = {
  regions: number;
  resources: number;
  accounts: number;
  cost: {
    date: string;
    amount: number;
  }[];
};

function useDashboardTopStats() {
  const [loading, setLoading] = useState(true);
  const [data, setData] = useState<Data>();
  const [error, setError] = useState(false);

  function fetch() {
    if (!loading) {
      setLoading(true);
    }

    if (error) {
      setError(false);
    }

    setTimeout(() => {
      setData(mockDataForDashboard.stats);
      setLoading(false);
    }, 1500);
  }

  useEffect(() => {
    fetch();
  }, []);

  return { loading, data, error, fetch };
}

export default useDashboardTopStats;
