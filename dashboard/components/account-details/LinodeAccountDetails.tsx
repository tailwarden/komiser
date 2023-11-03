import classNames from 'classnames';
import { LinodeCredentials } from '@utils/cloudAccountHelpers';
import LabelledInput from '../onboarding-wizard/LabelledInput';
import { CloudAccountPayload } from '../cloud-account/hooks/useCloudAccounts/useCloudAccount';

interface LinodeAccountDetailsProps {
  cloudAccountData?: CloudAccountPayload<LinodeCredentials>;
  hasError?: boolean;
}

function LinodeAccountDetails({
  cloudAccountData,
  hasError = false
}: LinodeAccountDetailsProps) {
  return (
    <div className="flex flex-col space-y-4 py-10">
      <LabelledInput
        type="text"
        id="account-name"
        name="name"
        value={cloudAccountData?.name}
        label="Account name"
        placeholder="my-linode-account"
      />

      <div
        className={classNames(
          'flex flex-col space-y-8 rounded-md p-5',
          hasError ? 'bg-red-50' : 'bg-gray-50'
        )}
      >
        <LabelledInput
          type="text"
          id="api-token"
          name="token"
          value={cloudAccountData?.credentials.token}
          label="API token"
          subLabel="An API key that is unique to your account"
          placeholder="abc123def456ghi789jkl012mno345pqr678stu901vwx234"
        />
      </div>
      {hasError && (
        <div className="text-sm text-red-500">
          We couldn&apos;t connect to your Linode account. Please check if the
          file is correct.
        </div>
      )}
    </div>
  );
}

export default LinodeAccountDetails;
