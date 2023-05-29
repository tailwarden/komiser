import { ReactNode } from 'react';

type OnboardingWizardLayoutProps = {
  children: ReactNode;
};

function OnboardingWizardLayout({ children }: OnboardingWizardLayoutProps) {
  return (
    <div className="flex flex-col gap-6">
      <p className="flex items-center gap-2 text-lg font-medium text-black-900">
        Onboarding Wizard
      </p>
      {children}
    </div>
  );
}

export default OnboardingWizardLayout;
