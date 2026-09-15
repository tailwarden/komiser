import classNames from 'classnames';
import RecordCircleIcon from '@components/icons/RecordCircleIcon';
import { HetznerCredentials } from '@utils/cloudAccountHelpers';
import LabelledInput from '../onboarding-wizard/LabelledInput';
import { CloudAccountPayload } from '../cloud-account/hooks/useCloudAccounts/useCloudAccount';

interface HetznerAccountDetailsProps {
  cloudAccountData?: CloudAccountPayload<HetznerCredentials>;
  hasError?: boolean;
}

function HetznerAccountDetails({
  cloudAccountData,
  hasError = false
}: HetznerAccountDetailsProps) {
  return (
    <div className="flex flex-col space-y-4 py-10">
      <LabelledInput
        type="text"
        id="account-name"
        name="name"
        value={cloudAccountData?.name}
        label="Account name"
        placeholder="my-hetzner-account"
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
          label="Cloud API token"
          subLabel="Create a read-only token at console.hetzner.cloud (Security > API Tokens)"
          placeholder="00000000000000000000000000000000"
        />
      </div>
      {hasError && (
        <div className="text-sm text-red-500">
          We couldn&apos;t connect to your Hetzner account. Please check if the
          token is correct.
        </div>
      )}
    </div>
  );
}

export default HetznerAccountDetails;
