import { useState } from 'react';
import { useRouter } from 'next/router';

import useToast from '../../../toast/hooks/useToast';
import { Provider } from '../../../../utils/providerHelper';

interface CloudAccounts {
  provider: Provider;
  name: string;
  status: {
    state: 'Connected' | 'Permission Issue' | 'Syncing';
    message: string;
  };
}

function useCloudAccount() {
  const router = useRouter();
  const { toast, dismissToast } = useToast();

  const [cloudAccounts, setCloudAccounts] = useState<Array<CloudAccounts>>([
    {
      provider: 'aws',
      name: 'Loudy AWS',
      status: {
        state: 'Connected',
        message: 'Your cloud account is connected.'
      }
    },
    {
      provider: 'azure',
      name: 'Loudy Azure',
      status: {
        state: 'Permission Issue',
        message:
          "We couldn't fetch EC2, S3 and VPC resources.See more details through the cloud account sidepanel."
      }
    },
    {
      provider: 'gcp',
      name: 'Loudy GCP',
      status: {
        state: 'Syncing',
        message: 'Your cloud account data is being fetched by Komiser.'
      }
    }
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
