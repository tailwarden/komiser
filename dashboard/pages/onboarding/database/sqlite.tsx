import Head from 'next/head';

import OnboardingWizardLayout, {
  LeftSideLayout,
  RightSideLayout
} from '../../../components/onboarding-wizard/OnboardingWizardLayout';
import LabelledInput from '../../../components/onboarding-wizard/LabelledInput';
import DatabasePurplin from '../../../components/onboarding-wizard/DatabasePurplin';
import CredentialsButton from '../../../components/onboarding-wizard/CredentialsButton';

export default function PostgreSQLCredentials() {
  const handleNext = () => {};

  return (
    <div>
      <Head>
        <title>Configure SQLite - Komiser</title>
        <meta name="description" content="Setup SQLite - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <OnboardingWizardLayout>
        <LeftSideLayout
          title="Configure your SQLite database"
          progressBarWidth="81%"
        >
          <div className="leading-6 text-gray-900/60">
            <div className="font-normal">
              SQLite is a lightweight, serverless, self-contained RDBMS that
              operates directly on files. It is known for its simplicity, ease
              of use, and portability across platforms.
            </div>
          </div>

          <div className="flex flex-col space-y-4 py-10">
            <div className="space-y-[0.2]">
              <LabelledInput
                type="text"
                id="file-path"
                label="File path"
                subLabel="Enter the path or browse the file"
                placeholder="C:\Documents\Komiser\database"
              />
            </div>
          </div>

          <CredentialsButton handleNext={handleNext} nextLabel="Add database" />
        </LeftSideLayout>

        <RightSideLayout>
          <DatabasePurplin database="sqlite" />
        </RightSideLayout>
      </OnboardingWizardLayout>
    </div>
  );
}
