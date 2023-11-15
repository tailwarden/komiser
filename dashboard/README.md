# 🚀 Komiser Dashboard

Komiser dashboard is a [Next.js](https://nextjs.org/) project bootstrapped with [`create-next-app`](https://github.com/vercel/next.js/tree/canary/packages/create-next-app).

**Full frontend stack:**

- 🖥 [`Next.js`](https://nextjs.org/)
- 📜 [`Typescript`](https://www.typescriptlang.org/)
- 🎨 [`Tailwind`](https://tailwindcss.com/)
- 📖 [`Storybook`](https://storybook.js.org/)
- 🧪 [`Jest`](https://jestjs.io/)
- 📚 [`React Testing Library`](https://testing-library.com/docs/react-testing-library/intro)

## 🚀 Getting Started

Follow the [Contribution Guide](https://github.com/tailwarden/komiser/blob/develop/CONTRIBUTING.md#contributing-to-komiser-dashboard-ui) first if you haven't done so already. Then come back here and follow the next steps:

#### 1. Run the development server:

From the Komiser root folder start the Komiser server by running:

```shell
go run *.go start --config /path/to/config.toml
```

In a different terminal tab navigate to the `/dashboard` folder:

```shell
cd dashboard
```

and run:

```shell
npm install

NEXT_PUBLIC_API_URL=http://localhost:3000 npm run dev
```

Alternatively, you can create an .env file with it, either manually or by running:

```shell
echo "NEXT_PUBLIC_API_URL=http://localhost:3000" > .env
```

and simply run:

```shell
npm run dev
```

#### 2. Open [http://localhost:3002/](http://localhost:3002). If you see the dashboard, 🎉 congrats! It's all up and running correctly.

❗ If you get an error page such as this, please refer to the logs and our [docs](https://docs.komiser.io/docs/introduction/getting-started).
<img alt="Error Image" src="https://user-images.githubusercontent.com/13384559/224320642-0bf6814b-d97a-4ad9-95a0-ca82e353c5d0.png" width="600"/>

## 🧩 Components

Komiser components are documented under `/components`

> 💡 **Hint:**
> We have the following import aliases defined in `tsconfig.json`
>
> ```json
> {
>   "@components/": "/dashboard/components/",
>   "@services/": "/dashboard/services/",
>   "@environments/": "/dashboard/environments/",
>   "@utils/": "/dashboard/utils/",
>   "@styles/": "/dashboard/styles/"
> }
> ```

You can find all the shared Components also inside [Storybook](https://storybook.komiser.io/). If you're implementing a new Story, please check for existing or new components with Storybook.
We will require a story for new shared components like icons, inputs or similar.

**Component convention:**

- 📁 Component folder: component name in `kebab-case`
- 📄 Component file: component name in `UpperCamelCase.*`
- 📖 Component story: component name in `UpperCamelCase.stories.*`
- 🎭 Component story mock (if needed): component name in `UpperCamelCase.mocks.*`
- 🧪 Component unit test: component name in `UpperCamelCase.test.*`
- 🧐 Check `Card` example for more details:

  <img alt="Component Example" src="https://user-images.githubusercontent.com/13384559/224307211-2ce62245-de24-4ee7-a156-fb54d8d34b4f.png" width="200"/>

**Additional instructions:**

- 📖 To view this component on Storybook, run: `npm run storybook`, then pick `Card`

  <img alt="Storybook Image" src="https://user-images.githubusercontent.com/13384559/224320112-e21d2ed4-1e22-4a33-adb3-6c236c4d4208.png" width="600"/>

- 🧪 To run the unit tests, run: `npm run test:watch`, hit `p`, then `card`

  <img alt="Unit Test Image" src="https://user-images.githubusercontent.com/13384559/224320260-19b1359e-1bfb-4db5-8379-918dacd7da44.png" width="400"/>

## 🧪 Testing

We use Jest & React Testing Library for our unit tests.

- To run the unit tests, run: `npm run test`

**Testing convention:**

- ✅ All new Utils need to be tested. Existing ones when being changed
- ✅ All tests should be wrapped in a `describe`
- ✅ If it's a unit test for a function: `describe('[replace with function name]', () => { ... })`
- ✅ If it's a unit test for a util: `describe('[replace with util name] util', () => { ... })`
- ✅ If it's a unit test for a component: `describe('[replace with component name]', () => { ... })`
- ✅ A test should use 'it' for the test function: `it('should do something', () => { ... })`

**Testing examples:**

- Simple Jest unit test example (snippet from `/utils/formatNumber.test.ts`):

```typescript
import formatNumber from './formatNumber';

describe('formatNumber util', () => {
  it('should format number (over a thousand) in short notation', () => {
    const result = formatNumber(12345);
    expect(result).toBe('12K');
  });
  ...
});
```

- Jest & Testing library example (snippet from `/components/card/Card.test.tsx`):

```typescript
import { render, screen } from '@testing-library/react';
import RefreshIcon from '../icons/RefreshIcon';
import Card from './Card';

describe('Card', () => {
  it('should render card component without crashing', () => {
    render(
      <Card
      label="Test card"
      value={500}
      icon={<RefreshIcon width={24} height={24} />}
      />
    );
  });

  it('should display the value formatted', () => {
    render(
      <Card
      label="Test card"
      value={5000}
      icon={<RefreshIcon width={24} height={24} />}
      />
    );
    const formattedNumber = screen.getByTestId('formattedNumber');
    expect(formattedNumber).toHaveTextContent('5K');
  });
  ...
});
```

If you're looking for an example with event firing and state updates, have a look at `components/select-checkbox/SelectCheckbox.test.tsx`:

```typescript
it('opens the dropdown when clicked', () => {
  const { getByRole, getByText } = render(
    <SelectCheckbox
      label="Test Label"
      query="provider"
      exclude={[]}
      setExclude={() => {}}
    />
  );

  fireEvent.click(getByRole('button'));

  expect(getByText('Item 1')).toBeInTheDocument();
  expect(getByText('Item 2')).toBeInTheDocument();
  expect(getByText('Item 3')).toBeInTheDocument();
});
```

## 🎨 Adding to Storybook

[**Storybook**](https://storybook.komiser.io/) is a tool for UI development. It makes development faster by isolating components. This allows you to work on one component at a time. If you create a new shared component or want to visualize variations of an existing one, follow these steps:

- To view this component on Storybook locally, run: `npm run storybook`, then pick an example (`Card`) or your new component story

  <img width="600" alt="image" src="https://user-images.githubusercontent.com/13384559/224320112-e21d2ed4-1e22-4a33-adb3-6c236c4d4208.png">

### 1. **Create the Story**:

In the same directory as your component, create a Storybook story:

- Create a story file: component name in `UpperCamelCase.stories.*`.

Here's a basic story format:

```typescript
import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';

import YourComponent from './YourComponent';

export default {
  title: 'Path/To/YourComponent',
  component: YourComponent
} as ComponentMeta<typeof YourComponent>;

const Template: ComponentStory<typeof YourComponent> = args => (
  <YourComponent {...args} />
);

export const Default = Template.bind({});
Default.args = {
  // default props here...
};
```

### 2. **Add Variations**:

You can create multiple variations of your component by replicating the `Default` pattern. For example, if your component has a variation for a "disabled" state:

```typescript
export const Disabled = Template.bind({});
Disabled.args = {
  // props to set the component to its disabled state...
};
```

### 3. **Mock Data**:

If your component requires mock data, create a mock file: component name in `UpperCamelCase.mocks.*`. Import this data into your story file to use with your component variations.

### 4. **Visual Check**:

Run Storybook:

```bash
npm run storybook
```

Your component should now appear in the Storybook UI. Navigate to it, and verify all the variations display correctly.

### 5. **Documentation**:

Add a brief description and any notes on your component's functionality within the Storybook UI. Use the `parameters` object in your default export:

```typescript
export default {
  title: 'Path/To/YourComponent',
  component: YourComponent,
  parameters: {
    docs: {
      description: {
        component: 'Your description here...'
      }
    }
  }
} as ComponentMeta<typeof YourComponent>;
```

---

> Remember: Storybook is not just a tool but also a way to document components. Ensure you provide meaningful names, descriptions, and use cases to help other developers understand the use and purpose of each component.

## 🤝 Contributing

We welcome all contributors to join us on the mission of improving Komiser, especially when it comes to writing tests and adding documentation.

Not sure where to start?

- 📖 Read the [contributor guidelines](https://docs.komiser.io/docs/introduction/community)
- 💬 [Join our Discord](https://discord.tailwarden.com/) and hang with us on #contributors channel.

## 📚 Learn More

To learn more about our stack, take a look at the following resources:

- [Next.js documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.
- [Tailwind documentation](https://tailwindcss.com/docs/)
- [Storybook documentation](https://storybook.js.org/docs/react/get-started/whats-a-story)
- [Jest documentation](https://jestjs.io/docs/getting-started)
- [React testing library documentation](https://testing-library.com/docs/dom-testing-library/intro)

## 🎥 Walkthrough video

[![Watch the video](https://komiser-assets-cdn.s3.eu-central-1.amazonaws.com/images/dashboard-contrib-video-thumb.png)](https://www.youtube.com/watch?v=uwxj11-eRt8)
