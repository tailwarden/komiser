import classNames from 'classnames';
import RecordCircleIcon from '@components/icons/RecordCircleIcon';
import { CivoCredentials } from '@utils/cloudAccountHelpers';
import LabelledInput from '../onboarding-wizard/LabelledInput';
import { CloudAccountPayload } from '../cloud-account/hooks/useCloudAccounts/useCloudAccount';

interface CivoAccountDetailsProps {
  cloudAccountData?: CloudAccountPayload<CivoCredentials>;
  hasError?: boolean;
}

function CivoAccountDetails({
  cloudAccountData,
  hasError = false
}: CivoAccountDetailsProps) {
  return (
    <div className="flex flex-col space-y-4 py-10">
      <LabelledInput
        type="text"
        id="account-name"
        name="name"
        value={cloudAccountData?.name}
        label="Account name"
        placeholder="my-civo-account"
      />

      <div
        className={classNames(
          'flex flex-col space-y-8 rounded-md p-5',
          hasError ? 'bg-red-50' : 'bg-gray-50'
        )}
      >
        <LabelledInput
          type="text"
          id="source"
          name="source"
          label="Source"
          value="API Token"
          disabled={true}
          icon={<RecordCircleIcon />}
        />
        <LabelledInput
          type="text"
          id="api-token"
          name="token"
          value={cloudAccountData?.credentials.token}
          label="API token"
          subLabel="An API key that is unique to your account"
          placeholder="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
        />
      </div>
      {hasError && (
        <div className="text-sm text-red-500">
          We couldn&apos;t connect to your Civo account. Please check if the
          file is correct.
        </div>
      )}
    </div>
  );
}

export default CivoAccountDetails;
