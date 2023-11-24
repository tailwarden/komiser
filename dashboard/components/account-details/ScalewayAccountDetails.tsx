import classNames from 'classnames';
import { ScalewayCredentials } from '@utils/cloudAccountHelpers';
import LabelledInput from '../onboarding-wizard/LabelledInput';
import { CloudAccountPayload } from '../cloud-account/hooks/useCloudAccounts/useCloudAccount';

interface ScalewayAccountDetailsProps {
  cloudAccountData?: CloudAccountPayload<ScalewayCredentials>;
  hasError?: boolean;
}

function ScalewayAccountDetails({
  cloudAccountData,
  hasError = false
}: ScalewayAccountDetailsProps) {
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
          id="access-key"
          name="accessKey"
          value={cloudAccountData?.credentials.accessKey}
          label="Access key"
          placeholder="SCWAKXXXXXXXXXXXXXXXXX"
        />
        <LabelledInput
          type="text"
          id="secret-key"
          name="secretKey"
          value={cloudAccountData?.credentials.secretKey}
          label="Secret key"
          placeholder="scwsk_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
        />
        <LabelledInput
          type="text"
          id="organization-id"
          name="organizationId"
          value={cloudAccountData?.credentials.organizationId}
          label="Organization ID"
          placeholder="demo"
        />
      </div>
      {hasError && (
        <div className="text-sm text-red-500">
          We couldn&apos;t connect to your Scaleway account. Please check if the
          file is correct.
        </div>
      )}
    </div>
  );
}

export default ScalewayAccountDetails;
