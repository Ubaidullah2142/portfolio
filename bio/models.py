from django.db import models

class Bio(models.Model):
    name = models.CharField(max_length=100)
    title = models.CharField(max_length=100)
    description = models.TextField()
    image = models.ImageField(upload_to='profile/')
    about_text = models.TextField()

    def __str__(self):
        return self.name
