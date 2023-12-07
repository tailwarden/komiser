import classNames from 'classnames';
import RecordCircleIcon from '@components/icons/RecordCircleIcon';
import { DigitalOceanCredentials } from '@utils/cloudAccountHelpers';
import LabelledInput from '../onboarding-wizard/LabelledInput';
import { CloudAccountPayload } from '../cloud-account/hooks/useCloudAccounts/useCloudAccount';

interface DigitalOceanAccountDetailsProps {
  cloudAccountData?: CloudAccountPayload<DigitalOceanCredentials>;
  hasError?: boolean;
}

function DigitalOceanAccountDetails({
  cloudAccountData,
  hasError = false
}: DigitalOceanAccountDetailsProps) {
  return (
    <div className="flex flex-col space-y-4 py-10">
      <LabelledInput
        type="text"
        id="account-name"
        name="name"
        value={cloudAccountData?.name}
        label="Account name"
        placeholder="my-digitalocean-account"
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
          value="Personal Access Token"
          disabled={true}
          icon={<RecordCircleIcon />}
        />
        <LabelledInput
          type="text"
          id="personal-access-token"
          name="token"
          label="Personal access token"
          subLabel="Personal access tokens function like ordinary OAuth access tokens"
          placeholder="abcd1234efgh5678ijklmnop90qrstuv"
        />
      </div>
      {hasError && (
        <div className="text-sm text-red-500">
          We couldn&apos;t connect to your Digital Ocean account. Please check
          if the file is correct.
        </div>
      )}
    </div>
  );
}

export default DigitalOceanAccountDetails;
