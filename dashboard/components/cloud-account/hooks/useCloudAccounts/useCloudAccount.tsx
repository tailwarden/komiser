import { useState } from 'react';
import { useRouter } from 'next/router';

import useToast from '../../../toast/hooks/useToast';
import { Provider } from '../../../../utils/providerHelper';

interface CloudAccounts {
  provider: Provider;
  name: string;
  status: 'Connected' | 'Permission Issue' | 'Syncing';
}

function useCloudAccount() {
  const router = useRouter();
  const { toast, dismissToast } = useToast();

  const [cloudAccounts, setCloudAccounts] = useState<Array<CloudAccounts>>([
    { provider: 'aws', name: 'Loudy AWS', status: 'Connected' },
    {
      provider: 'azure',
      name: 'Loudy Azure',
      status: 'Permission Issue'
    },
    { provider: 'gcp', name: 'Loudy GCP', status: 'Syncing' }
  ]);
  const isNotCustomView = !router.query.view;

  return {
    router,
    toast,
    dismissToast,
    cloudAccounts,
    isNotCustomView
  };
}

export default useCloudAccount;
