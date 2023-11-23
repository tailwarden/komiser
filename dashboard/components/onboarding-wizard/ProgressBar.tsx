function OnboardingWizardProgressBar({ width }: { width: string }) {
  return (
    <div className="sticky top-0 z-10 h-1 w-full bg-gray-100">
      <div className="h-1 bg-darkcyan-500" style={{ width }}></div>
    </div>
  );
}

export default OnboardingWizardProgressBar;
