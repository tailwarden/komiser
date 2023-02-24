import { useEffect, useState } from 'react';
import settingsService from '../../../services/settingsService';

function useSelectCheckbox(query: string) {
  const [listOfExcludableItems, setListOfExcludableItems] = useState<string[]>(
    []
  );
  const [error, setError] = useState(false);

  useEffect(() => {
    setError(false);

    if (query === 'provider') {
      settingsService.getProviders().then(res => {
        if (res === Error) {
          setError(true);
        } else {
          setListOfExcludableItems(res);
        }
      });
    }

    if (query === 'account') {
      settingsService.getAccounts().then(res => {
        if (res === Error) {
          setError(true);
        } else {
          setListOfExcludableItems(res);
        }
      });
    }

    if (query === 'region') {
      settingsService.getRegions().then(res => {
        if (res === Error) {
          setError(true);
        } else {
          setListOfExcludableItems(res);
        }
      });
    }

    if (query === 'service') {
      settingsService.getServices().then(res => {
        if (res === Error) {
          setError(true);
        } else {
          setListOfExcludableItems(res);
        }
      });
    }
  }, [query]);

  return { listOfExcludableItems, error };
}

export default useSelectCheckbox;
