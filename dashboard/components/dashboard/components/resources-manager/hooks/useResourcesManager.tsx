import { useEffect, useState } from 'react';
import settingsService from '../../../../../services/settingsService';

export type ResourcesManagerData = {
  label: string;
  total: number;
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

  function fetch(newQuery: ResourcesManagerQuery = 'region') {
    if (!loading) {
      setLoading(true);
    }

    if (error) {
      setError(false);
    }

    const payload = { filter: newQuery, exclude: [] };
    const payloadJson = JSON.stringify(payload);

    settingsService.getGlobalResources(payloadJson).then(res => {
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
    fetch(query);
  }, [query]);

  return { loading, data, error, fetch, query, setQuery };
}

export default useResourcesManager;
