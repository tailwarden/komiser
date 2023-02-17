import { useEffect, useState } from 'react';
import settingsService from '../../../../../services/settingsService';

export type DashboardCloudMapRegions = {
  name: string;
  label: string;
  latitude: string;
  longitude: string;
  resources: number;
}[];

function useCloudMap() {
  const [loading, setLoading] = useState(true);
  const [data, setData] = useState<DashboardCloudMapRegions>();
  const [error, setError] = useState(false);

  function fetch() {
    if (!loading) {
      setLoading(true);
    }

    if (error) {
      setError(false);
    }

    settingsService.getGlobalLocations().then(res => {
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
    fetch();
  }, []);

  return { loading, data, error, fetch };
}

export default useCloudMap;
