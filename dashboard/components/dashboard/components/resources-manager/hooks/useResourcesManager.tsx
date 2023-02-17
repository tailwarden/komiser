import { useEffect, useState } from 'react';
import mockDataForDashboard from '../../../utils/mockDataForDashboard';

export type ResourcesManagerData = {
  name: string;
  amount: number;
}[];

function useResourcesManager() {
  const [loading, setLoading] = useState(true);
  const [data, setData] = useState<ResourcesManagerData>();
  const [error, setError] = useState(false);

  function fetch() {
    if (!loading) {
      setLoading(true);
    }

    if (error) {
      setError(false);
    }

    setTimeout(() => {
      setData(mockDataForDashboard.resources);
      setLoading(false);
    }, 1500);
  }

  useEffect(() => {
    fetch();
  }, []);

  return { loading, data, error, fetch };
}

export default useResourcesManager;
