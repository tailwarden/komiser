import { useEffect, useState } from 'react';
import settingsService from '../../../services/settingsService';

type Data = {
  regions: number;
  resources: number;
  accounts: number;
  costs: number;
};

function useGlobalStats() {
  const [loading, setLoading] = useState(true);
  const [data, setData] = useState<Data>();
  const [error, setError] = useState(false);
  const [hasNoAccounts, setHasNoAccounts] = useState(false);

  function fetch() {
    if (!loading) {
      setLoading(true);
    }

    if (error) {
      setError(false);
    }

    settingsService.getGlobalStats().then(res => {
      if (res === Error) {
        setLoading(false);
        setError(true);
      } else {
        setLoading(false);
        // in the case of res.account === 0 the default view should be shown 
        // this couples the backend and the frontend strongly 
        // if (res.accounts === 0) {

        // } else {
        //   setData(res);
        // }
        setData(res);
      }
    });
  }

  useEffect(() => {
    fetch();
  }, []);

  return { loading, data, error, hasNoAccounts, fetch };
}

export default useGlobalStats;
