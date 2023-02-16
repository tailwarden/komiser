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
      const response = {
        resources: 522,
        regions: 17,
        costs: 680.908480745776,
        accounts: 0
      };
      if (res === Error) {
        setError(true);
      } else {
        setLoading(false);

        if (response.accounts === 0) {
          setHasNoAccounts(true);
        } else {
          setData(res);
        }
      }
    });
  }

  useEffect(() => {
    fetch();
  }, []);

  return { loading, data, error, hasNoAccounts, fetch };
}

export default useGlobalStats;
