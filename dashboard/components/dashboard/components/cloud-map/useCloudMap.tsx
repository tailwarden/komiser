import { useEffect, useState } from 'react';
import mockDataForDashboard from '../../utils/mockDataForDashboard';

export type Regions = {
  name: string;
  label: string;
  latitude: string;
  longitude: string;
  resources: number;
}[];

function useCloudMap() {
  const [loading, setLoading] = useState(true);
  const [data, setData] = useState<Regions>();
  const [error, setError] = useState(false);

  function fetch() {
    if (!loading) {
      setLoading(true);
    }

    if (error) {
      setError(false);
    }

    setTimeout(() => {
      setData(mockDataForDashboard.regions);
      setLoading(false);
    }, 1500);
  }

  useEffect(() => {
    fetch();
  }, []);

  return { loading, data, error, fetch };
}

export default useCloudMap;
