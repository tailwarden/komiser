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
  const [isLoading, setIsLoading] = useState(true);
  const [hasError, setHasError] = useState(false);

  useEffect(() => {
    if (!isLoading) {
      setIsLoading(true);
    }

    settingsService
      .getOnboardingStatus()
      .then(res => {
        if (
          res !== Error &&
          res.onboarded === false &&
          res.status === 'PENDING_DATABASE'
        ) {
          router.push('/onboarding/choose-database');
        } else {
          router.push('/onboarding/choose-cloud');
        }
      })
      .finally(() => setIsLoading(false));

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
    router,
    openModal,
    page,
    cloudAccountItem,
    setCloudAccountItem,
    goTo,
    toast,
    hasError,
    setToast,
    dismissToast,
    cloudAccounts,
    setCloudAccounts,
    isNotCustomView,
    isLoading
  };
}

export default useCloudAccount;
