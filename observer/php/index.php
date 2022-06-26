<?php

interface NotificationObservable {
    public function onMessage(string $message);
}

class NotificationSubject {
    private static $observers = [];

    public static function subscribe(NotificationObservable $observable): bool {
        self::$observers[] = $observable;
        return true;
    }
    
    public static function unsubscribe(NotificationObservable $observable): bool {
        foreach (self::$observers as $key => $obs) {
            if ($obs == $observable) {
                unset(self::$observers[$key]);
                return true;
            }
        }
        return false;
    }

    public static function notify(string $message) {
        foreach (self::$observers as $key => $obs) {
            $obs->onMessage($message);
        }
    }
}


class UserObserver implements NotificationObservable {
    private $id;
    public function __construct($id)
    {
        $this->id = $id;   
    }
    public function onMessage(string $message) {
        echo "OnMessage called from Object with ID: $this->id Message: $message" . PHP_EOL;
    }

}

$observer1 = new UserObserver(1);
$observer2 = new UserObserver(2);

NotificationSubject::subscribe($observer1);
NotificationSubject::subscribe($observer2);

NotificationSubject::notify('Message A');

NotificationSubject::unsubscribe($observer1);
NotificationSubject::notify('Message B');