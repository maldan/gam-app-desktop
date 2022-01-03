import { createStore, Store } from 'vuex';
import modal, { ModalStore } from '../gam_sdk_ui/vue/store/modal';
import { InjectionKey } from 'vue';

export interface MainTree {
  modal: ModalStore;
}

export const key: InjectionKey<Store<MainTree>> = Symbol();

export default createStore<MainTree>({
  modules: { modal },
});
