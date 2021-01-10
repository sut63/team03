import { createPlugin } from '@backstage/core';
import SignIn from './components/SignIn';
import MedicalfileCreate from './components/SaveMed';
import Menu from './components/Menu';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', SignIn);
    router.registerRoute('/Menu', Menu);
    router.registerRoute('/SaveMed', MedicalfileCreate);
    
  }
});
