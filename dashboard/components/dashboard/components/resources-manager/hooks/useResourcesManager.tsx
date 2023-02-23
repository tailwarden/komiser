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
  const [exclude, setExclude] = useState<string[]>([]);
  const [listOfResources, setListOfResources] = useState<string[]>([
    'AWS',
    'Kubernetes',
    'Civo',
    'Azure',
    'Other'
  ]);

  function fetch(filter: ResourcesManagerQuery = 'provider') {
    if (!loading) {
      setLoading(true);
    }

    if (error) {
      setError(false);
    }

    const payload = { filter, exclude };
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
  }, [query, exclude]);

  return {
    loading,
    data,
    error,
    fetch,
    query,
    setQuery,
    listOfResources,
    exclude,
    setExclude
  };
}

export default useResourcesManager;
