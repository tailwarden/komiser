import { useEffect, useState } from 'react';
import mockDataForDashboard from '../../../utils/mockDataForDashboard';

export type ResourcesManagerData = {
  name: string;
  amount: number;
}[];

export type ResourcesManagerQuery =
  | 'provider'
  | 'service'
  | 'region'
  | 'account'
  | 'view';

export type ResourcesManagerGroupBySelectProps = {
  values: ResourcesManagerQuery[];
  displayValues: string[];
};

function useResourcesManager() {
  const [loading, setLoading] = useState(true);
  const [data, setData] = useState<ResourcesManagerData>();
  const [error, setError] = useState(false);
  const [query, setQuery] = useState<ResourcesManagerQuery>('provider');

  function fetch(newQuery: ResourcesManagerQuery = 'provider') {
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
    fetch(query);
  }, [query]);

  return { loading, data, error, fetch, query, setQuery };
}

export default useResourcesManager;
