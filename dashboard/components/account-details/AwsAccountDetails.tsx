import { ChangeEvent, ReactNode, useRef, useState } from 'react';
import classNames from 'classnames';
import { AWSCredentials } from '@utils/cloudAccountHelpers';
import Folder2Icon from '../icons/Folder2Icon';
import SelectInput from '../onboarding-wizard/SelectInput';
import LabelledInput from '../onboarding-wizard/LabelledInput';
import InputFileSelect from '../onboarding-wizard/InputFileSelect';
import KeyIcon from '../icons/KeyIcon';
import VariableIcon from '../icons/VariableIcon';
import DocumentTextIcon from '../icons/DocumentTextIcon';
import ShieldSecurityIcon from '../icons/ShieldSecurityIcon';
import { CloudAccountPayload } from '../cloud-account/hooks/useCloudAccounts/useCloudAccount';

interface SelectOptions {
  icon: ReactNode;
  label: string;
  value: string;
}

interface AwsAccountDetailsProps {
  cloudAccountData?: CloudAccountPayload<AWSCredentials>;
  hasError?: boolean;
}

const options: SelectOptions[] = [
  {
    icon: <DocumentTextIcon />,
    label: 'Credentials File',
    value: 'credentials-file'
  },
  {
    icon: <KeyIcon />,
    label: 'Credentials keys',
    value: 'credentials-keys'
  },
  {
    icon: <VariableIcon />,
    label: 'Environment Variables',
    value: 'environment-variables'
  },
  {
    icon: <ShieldSecurityIcon />,
    label: 'IAM Instance Role',
    value: 'iam-instance-role'
  }
];

function AwsAccountDetails({
  cloudAccountData,
  hasError = false
}: AwsAccountDetailsProps) {
  const [credentialType, setCredentialType] = useState<string>(
    options.find(
      option => option.value === cloudAccountData?.credentials.source
    )?.value ?? options[0].value
  );
  const [isValidationError, setIsValidationError] = useState<boolean>(false);
  const [errorMessage, setErrorMessage] = useState<string>('');
  const [file, setFile] = useState<string>(
    cloudAccountData?.credentials.path || ''
  );

  const fileInputRef = useRef<HTMLInputElement | null>(null);
  const handleButtonClick = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  function handleSelectChange(newValue: string) {
    setCredentialType(newValue);
  }

  const handleFileChange = (event: ChangeEvent<HTMLInputElement>) => {
    const fileName = event.target.files?.[0]?.name;

    if (fileName) {
      setFile(fileName);
      if (!fileName.endsWith('.db')) {
        setIsValidationError(true);
        setErrorMessage(
          'The chosen file is not supported. Please choose a different file for the credentials.'
        );
      }
    } else {
      setIsValidationError(true);
      setErrorMessage('Please choose a file.');
    }
  };

  return (
    <div className="flex flex-col space-y-4 py-10">
      <LabelledInput
        type="text"
        id="account-name"
        name="name"
        label="Account name"
        placeholder="my-aws-account"
        required
        value={cloudAccountData?.name}
      />

      <div
        className={classNames(
          'flex flex-col space-y-8 rounded-md p-5',
          hasError ? 'bg-red-50' : 'bg-gray-50'
        )}
      >
        <div>
          <SelectInput
            icon="Change"
            name="source"
            label="Source"
            displayValues={options}
            value={credentialType}
            handleChange={handleSelectChange}
            values={options.map(option => option.value)}
          />
          {[options[2].value, options[3].value].includes(credentialType) && (
            <div className="mt-2 text-sm text-gray-700">
              {credentialType === options[3].value
                ? 'Komiser will fetch the credentials from AWS'
                : 'Komiser will load credentials from AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY.'}
            </div>
          )}
        </div>

        {credentialType === options[0].value && (
          <div>
            <InputFileSelect
              type="text"
              label="File path"
              id="file-path-input"
              name="path"
              icon={<Folder2Icon />}
              subLabel="Enter the path or browse the file"
              placeholder="C:\Documents\Komiser\credentials"
              fileInputRef={fileInputRef}
              iconClick={handleButtonClick}
              handleFileChange={handleFileChange}
              handleInputChange={e => setFile(e.target.value)}
              value={file}
              hasError={isValidationError}
              errorMessage={errorMessage}
            />
            <LabelledInput
              type="text"
              id="profile"
              name="profile"
              label="Profile"
              placeholder="default"
              subLabel="Name of the section in the credentials file"
              value={cloudAccountData?.credentials.profile}
              required
            />
          </div>
        )}

        {credentialType === options[1].value && (
          <div>
            <LabelledInput
              type="text"
              id="access-key-id"
              name="aws_access_key_id"
              label="Access key ID"
              placeholder="AKIABCDEFGHIJKLMN12"
              subLabel="Unique identifier used to access AWS services"
              required
            />
            <LabelledInput
              type="text"
              id="secret-access-key"
              name="aws_secret_access_key"
              label="Secret access key"
              placeholder="AbCdEfGhIjKlMnOpQrStUvWxYz0123456789AbCd"
              subLabel="The secret access key is generated by AWS when an access key is created"
              required
            />
          </div>
        )}
      </div>
      {hasError && (
        <div className="text-sm text-red-500">
          We couldn&apos;t connect to your AWS account. Please check if the file
          is correct.
        </div>
      )}
    </div>
  );
}

export default AwsAccountDetails;
