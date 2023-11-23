import { useState } from 'react';
import { ToastProps } from '@components/toast/Toast';
import AlertCircleIcon from '../../icons/AlertCircleIcon';
import Button from '../../button/Button';
import { CloudAccount } from '../hooks/useCloudAccounts/useCloudAccount';
import settingsService from '../../../services/settingsService';

interface CloudAccountDeleteContentsProps {
  cloudAccount: CloudAccount;
  onCancel: () => void;
  handleAfterDelete: (account: CloudAccount) => void;
  showToast: (toast: ToastProps) => void;
}

function CloudAccountDeleteContents({
  cloudAccount,
  onCancel,
  handleAfterDelete,
  showToast
}: CloudAccountDeleteContentsProps) {
  const [loading, setLoading] = useState(false);

  const deleteCloudAccount = () => {
    if (!cloudAccount.id) return false;

    setLoading(true);

    settingsService.deleteCloudAccount(cloudAccount.id).then(res => {
      setLoading(false);
      if (res === Error) {
        showToast({
          hasError: true,
          title: 'Cloud account was not deleted',
          message:
            'There was an error deleting this cloud account. Please try again.'
        });
      } else {
        showToast({
          hasError: false,
          title: 'Cloud account deleted',
          message: `The cloud account was successfully deleted!`
        });
        handleAfterDelete(cloudAccount);
      }
    });

    return true;
  };

  return (
    <>
      <div className="flex flex-col items-center gap-y-6">
        <AlertCircleIcon className="h-16 w-16" />
        <h1 className="text-center text-xl font-semibold text-gray-950">
          Are you sure you want to
          <br />
          remove this cloud account?
        </h1>
        <h3 className="text-center">
          All related data (like custom views and tags) will be deleted
          <br />
          and the {cloudAccount.name} account will be disconnected from Komiser.
        </h3>
      </div>
      <div className="flex flex-row place-content-end gap-x-8">
        <Button style="text" onClick={() => onCancel()}>
          Cancel
        </Button>
        <Button
          style="delete"
          loading={loading}
          onClick={() => deleteCloudAccount()}
        >
          Delete account
        </Button>
      </div>
    </>
  );
}

export default CloudAccountDeleteContents;
