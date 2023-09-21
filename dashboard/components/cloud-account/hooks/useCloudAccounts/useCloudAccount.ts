import { useEffect, useState } from 'react';
import { useRouter } from 'next/router';

import useToast from '../../../toast/hooks/useToast';
import { Provider } from '../../../../utils/providerHelper';
import settingsService from '../../../../services/settingsService';

export interface CloudAccount {
  credentials: {
    path: string;
    profile: string;
    source: string;
  };
  id?: number;
  name: string;
  provider: Provider;
  resources?: number;
  status?: 'CONNECTED' | 'INTEGRATION_ISSUE' | 'PERMISSION_ISSUE';
}

export type CloudAccountsPage = 'cloud account details';

function useCloudAccount() {
  const router = useRouter();
  const { toast, setToast, dismissToast } = useToast();
  const [page, setPage] = useState<CloudAccountsPage>('cloud account details');
  const [cloudAccounts, setCloudAccounts] = useState<CloudAccount[]>([]);
  const [cloudAccountItem, setCloudAccountItem] = useState<CloudAccount>();
  const isNotCustomView = !router.query.view;
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(false);

  useEffect(() => {
    settingsService.getCloudAccounts().then(res => {
      if (!loading) {
        setLoading(true);
      }

      if (error) {
        setError(false);
      }

      if (res === Error) {
        setLoading(false);
        setError(true);
      } else {
        setLoading(false);
        setCloudAccounts(res);
      }
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
    router,
    openModal,
    page,
    cloudAccountItem,
    setCloudAccountItem,
    goTo,
    toast,
    setToast,
    dismissToast,
    cloudAccounts,
    setCloudAccounts,
    isNotCustomView
  };
}

export default useCloudAccount;
