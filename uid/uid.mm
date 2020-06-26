//
//  uid.mm
//  uid
//
//  Created by mwpcheung on 30/5/20.
//  Copyright © 2020 mwpcheung. All rights reserved.
//

#import <Foundation/Foundation.h>

extern "C" bool NewUUID(char*uid)
{
        NSUUID* uuid = [NSUUID UUID];
    if (uid != nil){
        [uuid getUUIDBytes:(unsigned char*)uid];
        return true;
    }
    return false;
}

bool NewUUIDPlus(char*uid) // objective c++
{
    return NewUUID(uid);
}
