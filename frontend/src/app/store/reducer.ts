import { createReducer, on } from '@ngrx/store';
import * as IdentityActions from './actions';
import { IdentityObject } from './typedefs';

export interface IdentityState {
    identity: IdentityObject | null;
}

export const initialState: IdentityState = {
    identity: null,
};

export const identityReducer = createReducer(
    initialState,
    on(IdentityActions.SetIdentity, (state, action) => ({ identity: action.user_info })),
);
