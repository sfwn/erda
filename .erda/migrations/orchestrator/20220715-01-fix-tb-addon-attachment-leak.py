"""
Generated by Erda Migrator.
Please implement the function entry, and add it to the list entries.
"""


import django.db.models


class TbAddonAttachment(django.db.models.Model):
    """
    generated by erda-cli
    """

    id = django.db.models.BigIntegerField()
    app_id = django.db.models.CharField()
    instance_id = django.db.models.CharField()
    create_time = django.db.models.DateTimeField()
    update_time = django.db.models.DateTimeField()
    is_deleted = django.db.models.CharField()
    options = django.db.models.CharField()
    org_id = django.db.models.CharField()
    project_id = django.db.models.CharField()
    application_id = django.db.models.CharField()
    routing_instance_id = django.db.models.CharField()
    runtime_name = django.db.models.CharField()
    inside_addon = django.db.models.CharField()
    tenant_instance_id = django.db.models.CharField()
    mysql_account_id = django.db.models.CharField()
    previous_mysql_account_id = django.db.models.CharField()
    mysql_account_state = django.db.models.CharField()
    
    class Meta:
        db_table = "tb_addon_attachment"

def entry():
    """
    please implement this and add it to the list entries
    """

    attaches = TbAddonAttachment.objects.filter(is_deleted='S')
    attach_dic = {}
    for attachment in attaches:
        uni_id = attachment.routing_instance_id + attachment.app_id
        if uni_id in attach_dic.keys():
            continue
        attach_dic[uni_id] = attachment

    for attach in attach_dic.values():
        cs = TbAddonAttachment.objects.filter(routing_instance_id=attach.routing_instance_id, is_deleted='N', app_id=attach.app_id).count()
        if cs > 0:
            TbAddonAttachment.objects.filter(routing_instance_id=attach.routing_instance_id, is_deleted='S', app_id=attach.app_id).delete()
        else:
            latest = TbAddonAttachment.objects.filter(routing_instance_id=attach.routing_instance_id, is_deleted='S', app_id=attach.app_id).order_by('-id')
            newAttach = TbAddonAttachment.objects.get(id=latest[0].id)
            newAttach.is_deleted = 'N'
            newAttach.save()
            TbAddonAttachment.objects.filter(routing_instance_id=attach.routing_instance_id, is_deleted ='S', app_id = attach.app_id).delete()


entries: [callable] = [
    entry,
]
