import { ReactNode } from 'react';

import OnboardingWizardHeader from './PageHeaders';
import OnboardingWizardProgressBar from './ProgressBar';

function OnboardingWizardLayout({ children }: { children: ReactNode }) {
  return (
    <div className="font-['Noto Sans'] grid h-screen grid-cols-11">
      {children}
    </div>
  );
}

type LeftSideLayoutProps = {
  title: string;
  children: ReactNode;
  progressBarWidth: string;
};

function LeftSideLayout({
  title,
  children,
  progressBarWidth
}: LeftSideLayoutProps) {
  return (
    <div className="no-scrollbar col-span-6 overflow-y-scroll">
      <OnboardingWizardProgressBar width={progressBarWidth} />
      <div className="p-20">
        <OnboardingWizardHeader title={title} />
        {children}
      </div>
    </div>
  );
}

interface RightSideLayoutProps {
  children: ReactNode;
  isCustom?: boolean;
  customClasses?: string;
}

function RightSideLayout({
  isCustom,
  children,
  customClasses
}: RightSideLayoutProps) {
  return (
    <>
      {isCustom ? (
        <div className={`col-span-5 bg-gray-50 ${customClasses}`}>
          {children}
        </div>
      ) : (
        <div className="col-span-5 flex items-center justify-center bg-gray-50 p-7">
          {children}
        </div>
      )}
    </>
  );
}

export default OnboardingWizardLayout;
export { LeftSideLayout, RightSideLayout };
