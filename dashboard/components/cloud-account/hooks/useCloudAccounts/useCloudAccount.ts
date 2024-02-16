import { useEffect, useState } from 'react';
import { useRouter } from 'next/router';

import { Credentials } from '@utils/cloudAccountHelpers';

import { Provider } from '../../../../utils/providerHelper';
import settingsService from '../../../../services/settingsService';

export interface CloudAccount {
  credentials: Credentials;
  id?: number;
  name: string;
  provider: Provider;
  resources?: number;
  status?: 'CONNECTED' | 'INTEGRATION ISSUE' | 'PERMISSION ISSUE' | 'SCANNING';
}

export interface CloudAccountPayload<T extends Credentials> {
  name: string;
  provider: Provider;
  credentials: T;
}

export type CloudAccountsPage = 'cloud account details';

function useCloudAccount() {
  const router = useRouter();
  const [page, setPage] = useState<CloudAccountsPage>('cloud account details');
  const [cloudAccounts, setCloudAccounts] = useState<CloudAccount[]>([]);
  const [cloudAccountItem, setCloudAccountItem] = useState<CloudAccount>();
  const isNotCustomView = !router.query.view;
  const [isLoading, setIsLoading] = useState(true);
  const [hasError, setHasError] = useState(false);

  useEffect(() => {
    if (!isLoading) {
      setIsLoading(true);
    }

    settingsService.getCloudAccounts().then(res => {
      if (res === Error) {
        setHasError(true);
      } else {
        setCloudAccounts(res);
      }

      setIsLoading(false);
    });
  }, []);

  function openModal(CloudAccountItem: CloudAccount) {
    setCloudAccountItem(CloudAccountItem);
  }

  /** Handles the page change inside the modal */
  function goTo(newPage: CloudAccountsPage) {
    setPage(newPage);
  }

  return {
    openModal,
    page,
    cloudAccountItem,
    setCloudAccountItem,
    goTo,
    hasError,
    cloudAccounts,
    setCloudAccounts,
    isNotCustomView,
    isLoading
  };
}

export default useCloudAccount;
