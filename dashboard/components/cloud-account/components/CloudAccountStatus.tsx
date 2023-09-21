import classNames from 'classnames';
import { CloudAccount } from '../hooks/useCloudAccounts/useCloudAccount';

function CloudAccountStatus({ status }: { status: CloudAccount['status'] }) {
  if (!status) return null;

  return (
    <div
      className={classNames(
        'relative inline-block rounded-3xl px-2 py-1 text-sm',
        {
          'bg-green-200 text-green-600': status === 'CONNECTED',
          'bg-red-200 text-red-600':
            status === 'PERMISSION_ISSUE' || status === 'INTEGRATION_ISSUE'
        }
      )}
    >
      <span>{status.charAt(0) + status.slice(1).toLocaleLowerCase()}</span>
    </div>
  );
}

export default CloudAccountStatus;
