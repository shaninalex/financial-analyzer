import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { tap } from 'rxjs';
import { VerifiableAddress } from './typedefs';
import { MessagesService } from '../shared/services/messages.service';
import { Message } from '../shared/services/messages.service';


@Injectable()
export class IdentityEffects {

    constructor(
        private actions$: Actions,
        private messagesService: MessagesService
    ) { }

    setIdentity$ = createEffect(() => this.actions$.pipe(
        ofType('[identity] Set'),
        tap((data: any) => {
            for (let i = 0; i < data.user_info.identity.verifiable_addresses.length; i++) {
                let address: VerifiableAddress = data.user_info.identity.verifiable_addresses[i];
                if (!address.verified) {
                    const message: Message = {
                        type: "warning",
                        text: `Your email address (${address.value}) is not verified. Check your email box.`,
                    }
                    this.messagesService.message.next(message)
                }
            }
        }),
    ), { dispatch: false});

}
