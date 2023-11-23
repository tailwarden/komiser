import { ReactNode } from 'react';

type DashboardLayoutProps = {
  children: ReactNode;
};

function DashboardLayout({ children }: DashboardLayoutProps) {
  return (
    <div className="flex flex-col gap-6">
      <p className="flex items-center gap-2 text-lg font-medium text-gray-950">
        Dashboard overview
      </p>
      {children}
    </div>
  );
}

export default DashboardLayout;
