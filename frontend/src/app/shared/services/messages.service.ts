import { Injectable } from "@angular/core";
import { BehaviorSubject } from "rxjs";

export interface Message {
    type: string
    text: string
}

@Injectable({
    providedIn: 'root'
})
export class MessagesService {
    public message: BehaviorSubject<any | null> = new BehaviorSubject<any | null>(null);
    public messages: any[] = [];

    constructor() {
        this.message.subscribe({
            next: new_message => {
                if (new_message) this.messages.push(new_message);
            }
        })
    }
}