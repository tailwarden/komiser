import { useEffect, useState } from 'react';
import type { Meta, StoryObj } from '@storybook/react';
import Button from '@components/button/Button';
import Sidepanel from './Sidepanel';
import SidepanelHeader from './SidepanelHeader';
import SidepanelTabs from './SidepanelTabs';
import SidepanelPage from './SidepanelPage';
import SidepanelFooter from './SidepanelFooter';

type SidepanelProps = {
  title: string;
  subtitle?: string;
  href?: string;
  imgSrc?: string;
  imgAlt?: string;
  deleteLabel?: string;
  saveLabel?: string;
  tabs: string[];
};

function SidepanelWrapper({
  title,
  tabs,
  saveLabel,
  ...others
}: SidepanelProps) {
  const [isOpen, setIsOpen] = useState(false);
  const [page, setPage] = useState<string>('');

  useEffect(() => {
    if (tabs && tabs.length > 0) {
      setPage(tabs[0]);
    }
  }, []);

  const open = () => setIsOpen(true);
  const close = () => setIsOpen(false);
  const goTo = (newPage: string) => setPage(newPage);

  return (
    <div className="h-96">
      <Button onClick={open}>Open Sidepanel</Button>
      <Sidepanel isOpen={isOpen} closeModal={close}>
        <SidepanelHeader
          title={title}
          goBack={close}
          deleteAction={close}
          closeModal={close}
          {...others}
        />
        <SidepanelTabs goTo={goTo} page={page} tabs={tabs} />
        {tabs &&
          tabs.length > 0 &&
          tabs?.map((tab: string) => (
            <SidepanelPage key={tab} page={page} param={tab}>
              <p>
                Lorem ipsum, dolor sit amet consectetur adipisicing elit.
                Tempora et officia tenetur est, minima veritatis doloremque
                accusantium distinctio animi nulla reprehenderit quod asperiores
                similique illum perferendis, reiciendis, ipsam doloribus sit!
                Quidem amet veritatis ipsa omnis inventore architecto, assumenda
                ad vero cupiditate pariatur natus nisi corporis. Nobis voluptas
                vitae similique cupiditate deleniti. Inventore eos iusto porro
                perspiciatis fugiat nam sequi eligendi voluptas autem. Vitae
                beatae animi porro, fugiat eligendi hic cumque illum!
                Consectetur culpa obcaecati dolore praesentium harum. Provident
                nesciunt repudiandae eligendi quos, minima sed dolore veniam
                consequatur delectus! Optio ratione cum dolor eaque
                necessitatibus numquam maiores inventore asperiores quisquam
                quidem?
              </p>
            </SidepanelPage>
          ))}
        <SidepanelFooter
          saveLabel={saveLabel}
          saveAction={close}
          closeModal={close}
        />
      </Sidepanel>
    </div>
  );
}

const meta: Meta<typeof SidepanelWrapper> = {
  title: 'Komiser/SidePanel',
  component: SidepanelWrapper,
  tags: ['autodocs'],
  argTypes: {}
};

export default meta;
type Story = StoryObj<typeof SidepanelWrapper>;

export const Primary: Story = {
  args: {
    title: 'Resource name',
    tabs: ['tab 1', 'tab 2', 'tab 3'],
    subtitle: 'service',
    href: 'https://docs.komiser.io/',
    imgSrc: '/assets/img/providers/aws.png',
    imgAlt: 'komiser',
    deleteLabel: 'delete',
    saveLabel: 'Save changes'
  }
};

export const Secondary: Story = {
  args: {
    title: 'Create new policy',
    tabs: ['tab 1', 'tab 2', 'tab 3'],
    deleteLabel: 'delete',
    saveLabel: 'Save changes'
  }
};
