import classNames from 'classnames';
import { TencentCredentials } from '@utils/cloudAccountHelpers';
import LabelledInput from '../onboarding-wizard/LabelledInput';
import { CloudAccountPayload } from '../cloud-account/hooks/useCloudAccounts/useCloudAccount';

interface TencentAccountDetailsProps {
  cloudAccountData?: CloudAccountPayload<TencentCredentials>;
  hasError?: boolean;
}

function TencentAccountDetails({
  cloudAccountData,
  hasError = false
}: TencentAccountDetailsProps) {
  return (
    <div className="flex flex-col space-y-4 py-10">
      <LabelledInput
        type="text"
        id="account-name"
        name="name"
        value={cloudAccountData?.name}
        label="Account name"
        placeholder="my-tencent-account"
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
          placeholder="abcd1234efgh5678ijkl"
        />
      </div>
      {hasError && (
        <div className="text-sm text-red-500">
          We couldn&apos;t connect to your Tencent account. Please check if the
          file is correct.
        </div>
      )}
    </div>
  );
}

export default TencentAccountDetails;
