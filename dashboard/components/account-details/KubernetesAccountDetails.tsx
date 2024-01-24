import { ReactNode, useRef, useState, ChangeEvent } from 'react';
import classNames from 'classnames';
import DocumentTextIcon from '@components/icons/DocumentTextIcon';
import SelectInput from '@components/onboarding-wizard/SelectInput';
import InputFileSelect from '@components/onboarding-wizard/InputFileSelect';
import Folder2Icon from '@components/icons/Folder2Icon';
import { KubernetesCredentials } from '@utils/cloudAccountHelpers';
import LabelledInput from '../onboarding-wizard/LabelledInput';
import { CloudAccountPayload } from '../cloud-account/hooks/useCloudAccounts/useCloudAccount';

interface KubernetesAccountDetailsProps {
  cloudAccountData?: CloudAccountPayload<KubernetesCredentials>;
  hasError?: boolean;
}

interface SelectOptions {
  icon: ReactNode;
  label: string;
  value: string;
}

const options: SelectOptions[] = [
  {
    icon: <DocumentTextIcon />,
    label: 'Credentials File',
    value: 'credentials-file'
  }
];

function KubernetesAccountDetails({
  cloudAccountData,
  hasError = false
}: KubernetesAccountDetailsProps) {
  const [credentialType, setCredentialType] = useState<string>(
    options.find(
      option => option.value === cloudAccountData?.credentials.source
    )?.value ?? options[0].value
  );
  const [isValidationError, setIsValidationError] = useState<boolean>(false);
  const [errorMessage, setErrorMessage] = useState<string>('');
  const [file, setFile] = useState<string>(
    cloudAccountData?.credentials.file || ''
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
      if (!fileName.endsWith('.db')) {
        setIsValidationError(true);
        setErrorMessage(
          'The chosen file is not supported. Please choose a different file for the credentials.'
        );
        setFile(fileName);
      }
    } else {
      setIsValidationError(true);
      setErrorMessage('Please choose a file.');
    }
  };

  return (
    <div className="flex flex-col space-y-8 py-10">
      <LabelledInput
        type="text"
        id="account-name"
        name="name"
        value={cloudAccountData?.name}
        label="Account name"
        placeholder="my-kubernetes-account"
      />
      <div
        className={classNames(
          'flex flex-col space-y-8 rounded-md p-5',
          hasError ? 'bg-red-50' : 'bg-gray-50'
        )}
      >
        <SelectInput
          icon="Change"
          label="Source"
          name="source"
          displayValues={options}
          value={credentialType}
          handleChange={handleSelectChange}
          values={options.map(option => option.value)}
        />
        <InputFileSelect
          type="text"
          id="file-path-input"
          name="file"
          label="File path"
          value={file}
          subLabel="Enter the path or browse the file"
          placeholder="C:\Documents\Komiser\credentials"
          icon={<Folder2Icon className="h-6 w-6" />}
          fileInputRef={fileInputRef}
          iconClick={handleButtonClick}
          handleFileChange={handleFileChange}
          handleInputChange={e => setFile(e.target.value)}
          hasError={isValidationError}
          errorMessage={errorMessage}
        />
        <LabelledInput
          type="text"
          id="opencost-base-url"
          name="opencostBaseUrl"
          value={cloudAccountData?.name}
          label="Opencost Base URL"
          placeholder="localhost:9003"
        />
      </div>
      {hasError && (
        <div className="text-sm text-red-500">
          We couldn&apos;t connect to your Kubernetes account. Please check if
          the file is correct.
        </div>
      )}
    </div>
  );
}

export default KubernetesAccountDetails;
