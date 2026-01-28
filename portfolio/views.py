from django.shortcuts import render, redirect
from bio.models import Bio
from skills.models import Skill
from services.models import Service
from projects.models import Project
from contact.models import Contact

def home(request):
    if request.method == 'POST':
        Contact.objects.create(
            name=request.POST['name'],
            email=request.POST['email'],
            message=request.POST['message']
        )
        return redirect('/')

    context = {
        'bio': Bio.objects.first(),
        'skills': Skill.objects.all(),
        'services': Service.objects.all(),
        'projects': Project.objects.all(),
    }
    return render(request, 'index.html', context)
