import { createPlugin } from '@backstage/core';
import SignIn from './components/SignIn';
import MedicalfileCreate from './components/SaveMed';
import Menu from './components/Menu';
import Patient from './components/SavePatient';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', SignIn);
    router.registerRoute('/Menu', Menu);
    router.registerRoute('/SaveMed', MedicalfileCreate);
    router.registerRoute('/SavePatient', Patient);
    
  }
});
