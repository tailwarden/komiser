import classNames from 'classnames';
import { MongoDBAtlasCredentials } from '@utils/cloudAccountHelpers';
import LabelledInput from '../onboarding-wizard/LabelledInput';
import { CloudAccountPayload } from '../cloud-account/hooks/useCloudAccounts/useCloudAccount';

interface MongoDbAtlasAccountDetailsProps {
  cloudAccountData?: CloudAccountPayload<MongoDBAtlasCredentials>;
  hasError?: boolean;
}

function MongoDbAtlasAccountDetails({
  cloudAccountData,
  hasError = false
}: MongoDbAtlasAccountDetailsProps) {
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
          hasError ? 'bg-error-100' : 'bg-komiser-100'
        )}
      >
        <LabelledInput
          type="text"
          id="public-key"
          name="publicApiKey"
          value={cloudAccountData?.credentials.publicApiKey}
          label="Public key"
          placeholder="ABCDWXYZ"
        />
        <LabelledInput
          type="text"
          id="private-key"
          name="privateApiKey"
          value={cloudAccountData?.credentials.privateApiKey}
          label="Private key"
          placeholder="abcdefgh12345678ijklmnop"
        />
        <LabelledInput
          type="text"
          id="organization-id"
          name="organizationId"
          value={cloudAccountData?.credentials.organizationId}
          label="Organization ID"
          placeholder="5d31e955ff7a25d2e5a7xxxx"
        />
      </div>
      {hasError && (
        <div className="text-sm text-error-600">
          We couldn&apos;t connect to your MongoDB Atlas account. Please check
          if the file is correct.
        </div>
      )}
    </div>
  );
}

export default MongoDbAtlasAccountDetails;
