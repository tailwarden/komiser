import { ReactNode } from 'react';

type SidepanelPageProps = {
  page: string;
  param: string;
  children: ReactNode;
  container?: boolean;
};

function SidepanelPage({
  page,
  param,
  children,
  container = false
}: SidepanelPageProps) {
  return (
    <>
      {page === param && (
        <div className={container ? 'rounded-lg bg-black-100 p-6' : ''}>
          <div className="flex flex-col gap-6">{children}</div>
        </div>
      )}
    </>
  );
}

export default SidepanelPage;
